package main

import (
	"log"
	"net"

	"github.com/vnotes/workweixin/services/cores/grpc/todo"
	"github.com/vnotes/workweixin/services/todosrv/apis"
	_ "github.com/vnotes/workweixin/services/todosrv/dbs"

	"google.golang.org/grpc"
)

const (
	rpcPort = ":11112"
)

func main() {
	log.Print("listening rpc port 11112")

	lis, err := net.Listen("tcp", rpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	todo.RegisterToDoCmdServer(s, &apis.Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
