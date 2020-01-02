package todo_grpc

import (
	"context"
	"encoding/json"

	"github.com/vnotes/workweixin_app/cores/todos"
)

func ToDoCmdRPC(ctx context.Context, cmd, userID, content string) (meta string) {
	req := &ToDoRequest{
		UserID:  userID,
		Content: content,
		ToDoID:  0,
	}
	switch cmd {
	case todos.ToDoGet:
		rsp, err := ToDoCli.Select(ctx, req)
		if err != nil {
			meta = "查询 TODO 任务发生错误：" + err.Error()
			return
		}
		data := rsp.MapResultList()
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
