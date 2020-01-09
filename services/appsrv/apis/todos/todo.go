package todos

import (
	"context"
	"encoding/json"
	"log"
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

func queryToDoList(ctx context.Context, req *todo.ToDoRequest) (string, error) {
	// 从缓存查询
	var isFound bool
	val, err := getToDoListByCache()
	if err == nil && val == "" {
		log.Println("从缓存查找 ToDoList 成功")
		isFound = true
	}
	if isFound {
		return val, nil
	}

	// 从数据库查询
	rsp, err := ToDoCli.Select(ctx, req)
	if err != nil {
		return "", err
	}
	data := foldToResultList(rsp)

	meta, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}

	// 缓存数据
	result := string(meta)
	if err := CacheToDoList(result); err != nil {
		log.Printf("cache todo list error %#v", err)
	}
	return result, nil
}

func ToDoCmd(ctx context.Context, cmd, userID, content string) (meta string) {
	req := &todo.ToDoRequest{
		UserID:  userID,
		Content: content,
		ToDoID:  0,
	}
	switch cmd {
	case ToDoGet:
		data, err := queryToDoList(ctx, req)
		if err != nil {
			log.Printf("query todo list error %#v", err)
			meta = "查询ToDoList失败 " + err.Error()
			return meta
		}
		meta = "查询 TODO LIST 成功\n" + data
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
