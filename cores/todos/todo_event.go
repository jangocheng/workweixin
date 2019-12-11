package todos

import (
	"context"
	_ "database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

type Todo interface {
	Get(ctx context.Context) ([]*ToDoList, error)
	Update(ctx context.Context)
	Add(ctx context.Context, name []string) error
	Delete(ctx context.Context)
}

type ToDoList struct {
	Name       string `db:"todo_name"`
	CreateTime string `db:"create_time"`
	FinishTime string `db:"finish_time"`
	Active     bool   `db:"active"`
}

type ToDoEvent struct {
	db *sqlx.DB
}

func (t *ToDoEvent) Get(ctx context.Context) ([]*ToDoList, error) {
	result := make([]*ToDoList, 0)
	querySQL := "select todo_name, create_time, finish_time, active from todo_list;"
	if err := t.db.SelectContext(ctx, &result, querySQL); err != nil {
		log.Printf("query from db error %#v", err)
		return nil, err
	}
	return result, nil
}

func (t *ToDoEvent) Update(ctx context.Context) {

}

func (t *ToDoEvent) Add(ctx context.Context, name []string) error {
	_sql := "insert into todo_list(todo_nam) values(?);"
	_, err := t.db.ExecContext(ctx, _sql, name)
	if err != nil {
		return err
	}
	return nil

}
func (t *ToDoEvent) Delete(ctx context.Context) {

}
