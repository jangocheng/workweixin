package todos

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/vnotes/workweixin/services/appsrv/conf"
	"github.com/vnotes/workweixin/services/cores/grpc/todo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	ToDoGet    = "todo:get"
	ToDoADD    = "todo:add"
	ToDoDone   = "todo:done"
	ToDoDEL    = "todo:del"
	ToDoUpdate = "todo:update"
)

const HELP = `
	-------- todo 用法 ----------
	通过查询 todo_list 获取 todo_id
	------------------------------
	查询todo: todo:get@
	新建todo: todo:add@明年我要增高
	完成todo: todo:done@<todo_id>
	删除todo: todo:del@<todo_id>
	更新todo: todo:update@<todo_id>|text
`

var (
	ToDoCli  todo.ToDoCmdClient
	ToDoConn *grpc.ClientConn
)

type MetaDataReaderWriter struct {
	metadata.MD
}

func (w MetaDataReaderWriter) Set(key, val string) {
	key = strings.ToLower(key)
	w.MD[key] = append(w.MD[key], val)
}

func (w MetaDataReaderWriter) ForeachKey(handler func(key, val string) error) error {
	for k, val := range w.MD {
		for _, v := range val {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func clientInterceptor(tracer opentracing.Tracer) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		var parentCtx opentracing.SpanContext
		if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
			parentCtx = parentSpan.Context()
		}
		childSpan := tracer.StartSpan(
			method,
			opentracing.ChildOf(parentCtx),
			opentracing.Tag{
				Key:   string(ext.Component),
				Value: "gRPC",
			},
			ext.SpanKindRPCClient,
		)
		defer childSpan.Finish()
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		var (
			newMD = MetaDataReaderWriter{md}
			err   error
		)

		err = tracer.Inject(
			childSpan.Context(),
			opentracing.TextMap,
			newMD,
		)
		if err != nil {
			log.Printf("inject metadata error %#v", err)
			return err
		}
		ctx = metadata.NewOutgoingContext(ctx, md)
		err = invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			log.Printf("invoker error %#v", err)
			return err
		}
		return nil
	}
}

func InitToDoGRPC(tracer opentracing.Tracer) {
	address := fmt.Sprintf("%s:11112", conf.Conf.ToDoNetWork)
	var (
		err    error
		ctx, _ = context.WithTimeout(context.Background(), time.Second*10)
	)

	ToDoConn, err = grpc.DialContext(ctx, address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(clientInterceptor(tracer)),
	)
	if err != nil {
		log.Fatalf("don't connect: %v", err)
	}
	ToDoCli = todo.NewToDoCmdClient(ToDoConn)
}
