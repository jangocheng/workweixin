package todos

import (
	"context"
	"encoding/json"

	"github.com/vnotes/workweixin/services/cores/grpc/todo"
)

type ToDoList struct {
	ID         uint64 `db:"id"`
	Name       string `db:"todo_name"`
	CreateTime string `db:"create_time"`
	FinishTime string `db:"finish_time"`
	Active     bool   `db:"active"`
}

func foldToResultList(response *todo.ToDoResponse) []*ToDoList {
	rs := response.Result
	var rsp []*ToDoList
	for _, k := range rs {
		rsp = append(rsp, &ToDoList{
			ID:         k.ID,
			Name:       k.Name,
			CreateTime: k.CreateTime,
			FinishTime: k.FinishTime,
			Active:     k.Active,
		})
	}
	return rsp
}
func ToDoCmd(ctx context.Context, cmd, userID, content string) (meta string) {
	req := &todo.ToDoRequest{
		UserID:  userID,
		Content: content,
		ToDoID:  0,
	}
	switch cmd {
	case ToDoGet:
		rsp, err := ToDoCli.Select(ctx, req)
		if err != nil {
			meta = "查询 TODO 任务发生错误：" + err.Error()
			return
		}

		data := foldToResultList(rsp)

		metaStr, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			meta = "序列化数据出错： " + err.Error()
			return
		}
		meta = "查询 TODO LIST 成功\n" + string(metaStr)
		return
	default:
		meta = "命名格式有问题，请输入 HELP 查看帮助"
	}
	return meta
}
