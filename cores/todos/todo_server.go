package todos

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	rpcPort = ":11112"
)

type server struct {
	UnimplementedToDoCmdServer
}

func (s *server) Select(ctx context.Context, in *ToDoRequest) (*ToDoResponse, error) {
	todoList := make([]*ToDoList, 0)
	querySQL := "SELECT id, todo_name, create_time, finish_time, active FROM todo_list WHERE user_id = ?;"
	if err := t.db.SelectContext(ctx, &todoList, querySQL, in.UserID); err != nil {
		return nil, err
	}
	var result []*ToDoResult
	for _, v := range todoList {
		result = append(result, &ToDoResult{
			ID:         v.ID,
			Name:       v.Name,
			CreateTime: v.CreateTime,
			FinishTime: v.FinishTime,
			Active:     v.Active,
		})
	}
	rsp := &ToDoResponse{
		Result: result,
	}
	return rsp, nil
}

func InitGRPCServer() {
	log.Print("listening rpc port 11112")
	lis, err := net.Listen("tcp", rpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterToDoCmdServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}