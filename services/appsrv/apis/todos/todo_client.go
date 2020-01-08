package todos

import (
	"fmt"
	"log"

	"github.com/vnotes/workweixin/services/appsrv/conf"
	"github.com/vnotes/workweixin/services/cores/grpc/todo"

	"google.golang.org/grpc"
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

func InitToDoGRPC() {
	address := fmt.Sprintf("%s:11112", conf.Conf.ToDoNetWork)
	var err error
	ToDoConn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("don't connect: %v", err)
	}
	ToDoCli = todo.NewToDoCmdClient(ToDoConn)
}
