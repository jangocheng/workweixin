package todo_grpc

import (
	"log"

	"google.golang.org/grpc"
)

var (
	ToDoCli  ToDoCmdClient
	ToDoConn *grpc.ClientConn
)

const (
	address = "localhost:11112"
)

func InitGRPCClient() {
	ToDoConn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("don't connect: %v", err)
	}
	ToDoCli = NewToDoCmdClient(ToDoConn)
}

func CloseGRPCClient() {
	if ToDoConn != nil {
		_ = ToDoConn.Close()
	}
}
