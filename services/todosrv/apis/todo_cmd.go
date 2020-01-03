package apis

import (
	"context"

	"github.com/vnotes/workweixin/services/cores/grpc/todo"
	"github.com/vnotes/workweixin/services/todosrv/dbs"
)

type Server struct {
	todo.UnimplementedToDoCmdServer
}

func (s *Server) Select(ctx context.Context, in *todo.ToDoRequest) (*todo.ToDoResponse, error) {
	todoList := make([]*ToDoList, 0)
	querySQL := "SELECT id, todo_name, create_time, finish_time, active FROM todo_list WHERE user_id = ?;"
	if err := dbs.Cli().SelectContext(ctx, &todoList, querySQL, in.UserID); err != nil {
		return nil, err
	}
	var result []*todo.ToDoResult
	for _, v := range todoList {
		result = append(result, &todo.ToDoResult{
			ID:         v.ID,
			Name:       v.Name,
			CreateTime: v.CreateTime,
			FinishTime: v.FinishTime,
			Active:     v.Active,
		})
	}
	rsp := &todo.ToDoResponse{
		Result: result,
	}
	return rsp, nil
}
