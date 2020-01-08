package apis

type ToDoList struct {
	ID         uint64 `db:"id"`
	Name       string `db:"todo_name"`
	CreateTime string `db:"create_time"`
	FinishTime string `db:"finish_time"`
	Active     bool   `db:"active"`
}
