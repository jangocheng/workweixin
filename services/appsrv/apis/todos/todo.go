package todos

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

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
	case ToDoUpdate:
		text := strings.Split(content, "|")
		if len(text) != 2 {
			meta = "命令出错，请输入 HELP"
			return
		}
		todoID, err := strconv.ParseUint(text[0], 10, 64)
		if err != nil {
			meta = "命令有误，请输入 HELP"
			return
		}
		req.ToDoID = todoID
		req.Content = text[1]
		_, err = ToDoCli.Update(ctx, req)
		if err != nil {
			meta = "更新TODO任务失败：" + err.Error()
			return
		}
		meta = "更新TODO任务成功"
		return
	case ToDoDone:
		todoID, err := strconv.ParseUint(content, 10, 64)
		if err != nil {
			meta = "命令有误，请输入 HELP"
			return
		}
		req.ToDoID = todoID
		_, err = ToDoCli.Done(ctx, req)
		if err != nil {
			meta = "完成TODO任务失败：" + err.Error()
			return
		}
		meta = "已经完成TODO任务"
		return
	case ToDoDEL:
		todoID, err := strconv.ParseUint(content, 10, 64)
		if err != nil {
			meta = "命令有误，请输入 HELP"
			return
		}
		req.ToDoID = todoID
		_, err = ToDoCli.Delete(ctx, req)
		if err != nil {
			meta = "删除TODO任务失败：" + err.Error()
			return
		}
		meta = "删除TODO任务成功"
		return
	case ToDoADD:
		_, err := ToDoCli.Create(ctx, req)
		if err != nil {
			meta = "添加TODO任务失败：" + err.Error()
			return
		}
		meta = "添加TODO任务成功"
		return
	default:
		meta = "命名格式有问题，请输入 HELP 查看帮助"
	}
	return meta
}
