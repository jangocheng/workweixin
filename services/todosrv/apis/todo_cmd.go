package apis

import (
	"context"
	"log"

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

func (s *Server) Update(ctx context.Context, in *todo.ToDoRequest) (*todo.Empty, error) {
	_sql := "UPDATE todo_list SET todo_name = ? WHERE user_id = ? AND id = ?;"
	_, err := dbs.Cli().ExecContext(ctx, _sql, in.Content, in.UserID, in.ToDoID)
	if err != nil {
		log.Printf("update db error %#v", err)
		return nil, err
	}
	return &todo.Empty{}, nil
}

func (s *Server) Create(ctx context.Context, in *todo.ToDoRequest) (*todo.Empty, error) {
	_sql := "INSERT INTO todo_list (user_id, todo_name) values(?, ?);"
	_, err := dbs.Cli().ExecContext(ctx, _sql, in.UserID, in.Content)
	if err != nil {
		return nil, err
	}
	return &todo.Empty{}, nil
}

func (s *Server) Delete(ctx context.Context, in *todo.ToDoRequest) (*todo.Empty, error) {
	_sql := "DELETE FROM todo_list WHERE user_id = ? AND id = ?;"
	_, err := dbs.Cli().ExecContext(ctx, _sql, in.UserID, in.ToDoID)
	if err != nil {
		return nil, err
	}
	return &todo.Empty{}, nil
}

func (s *Server) Done(ctx context.Context, in *todo.ToDoRequest) (*todo.Empty, error) {
	_sql := "UPDATE todo_list SET active = 1 WHERE user_id = ? AND id = ?;"
	_, err := dbs.Cli().ExecContext(ctx, _sql, in.UserID, in.ToDoID)
	if err != nil {
		return nil, err
	}
	return &todo.Empty{}, nil
}
