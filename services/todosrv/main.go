package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/vnotes/workweixin/services/cores/grpc/todo"
	"github.com/vnotes/workweixin/services/todosrv/apis"
	"github.com/vnotes/workweixin/services/todosrv/conf"
	"github.com/vnotes/workweixin/services/todosrv/dbs"
	"github.com/vnotes/workweixin/services/todosrv/tracings"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
)

const (
	rpcPort     = ":11112"
	ServiceName = "weixin.todosrv"
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

func serverInterceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}
		spanContext, err := tracer.Extract(opentracing.TextMap, MetaDataReaderWriter{md})
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			grpclog.Errorf("extract error %#v", err)
			return handler(ctx, req)
		}
		serverSpan := tracer.StartSpan(
			info.FullMethod,
			ext.RPCServerOption(spanContext),
			opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
			ext.SpanKindRPCServer,
		)
		defer serverSpan.Finish()
		ctx = opentracing.ContextWithSpan(ctx, serverSpan)
		return handler(ctx, req)
	}
}

func main() {
	tracings.InitTracing(ServiceName)
	defer func() {
		_ = tracings.CloseTracer()
	}()

	conf.InitConfig()
	dbs.InitMySQL()

	log.Print("listening rpc port 11112")

	lis, err := net.Listen("tcp", rpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(serverInterceptor(tracings.Tracer)))
	todo.RegisterToDoCmdServer(s, &apis.Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
