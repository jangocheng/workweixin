package todos

import (
	"context"
	"log"
	"sync"

	"github.com/vnotes/workweixin_app/cores/dbs"

	"github.com/jmoiron/sqlx"
)

type Todo interface {
	Get(ctx context.Context, userID string) ([]*ToDoList, error)
	Update(ctx context.Context, userID, content string, todoID uint64) error
	Create(ctx context.Context, userID, content string) error
	Delete(ctx context.Context, userID string, toDoID uint64) error
	Done(ctx context.Context, userID string, toDoID uint64) error
}

var (
	once sync.Once

	t *ToDoEvent
)

func Cli() Todo {
	once.Do(func() {
		if t == nil {
			t = &ToDoEvent{db: dbs.DB}
		}
	})
	return t
}

type ToDoList struct {
	ID         uint64 `db:"id"`
	Name       string `db:"todo_name"`
	CreateTime string `db:"create_time"`
	FinishTime string `db:"finish_time"`
	Active     bool   `db:"active"`
}

type ToDoEvent struct {
	db *sqlx.DB
}

func (t *ToDoEvent) Get(ctx context.Context, userID string) ([]*ToDoList, error) {
	result := make([]*ToDoList, 0)
	querySQL := "SELECT id, todo_name, create_time, finish_time, active FROM todo_list WHERE user_id = ?;"
	if err := t.db.SelectContext(ctx, &result, querySQL, userID); err != nil {
		return nil, err
	}
	return result, nil
}

func (t *ToDoEvent) Update(ctx context.Context, userID, content string, todoID uint64) error {
	_sql := "UPDATE todo_list SET todo_name = ? WHERE user_id = ? AND id = ?;"
	_, err := t.db.ExecContext(ctx, _sql, content, userID, todoID)
	if err != nil {
		log.Printf("update db error %#v", err)
		return err
	}
	return nil
}

func (t *ToDoEvent) Create(ctx context.Context, userID, name string) error {
	_sql := "INSERT INTO todo_list (user_id, todo_name) values(?, ?);"
	_, err := t.db.ExecContext(ctx, _sql, userID, name)
	if err != nil {
		return err
	}
	return nil
}

func (t *ToDoEvent) Delete(ctx context.Context, userID string, todoID uint64) error {
	_sql := "DELETE FROM todo_list WHERE user_id = ? AND id = ?;"
	_, err := t.db.ExecContext(ctx, _sql, userID, todoID)
	if err != nil {
		return err
	}
	return nil
}

func (t *ToDoEvent) Done(ctx context.Context, userID string, todoID uint64) error {
	_sql := "UPDATE todo_list SET active = 1 WHERE user_id = ? AND id = ?;"
	_, err := t.db.ExecContext(ctx, _sql, userID, todoID)
	if err != nil {
		return err
	}
	return nil
}
