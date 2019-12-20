package todos

import (
	"context"
	"encoding/json"
)

const (
	ToDoGet    = "todo:get"
	ToDoADD    = "todo:add"
	ToDoDone   = "todo:done"
	ToDoDEL    = "todo:del"
	ToDoUpdate = "todo:update"
)

const HELP = `
	todo 用法----- \n
	通过查询 todo_list 获取 todo_id \n
	------------------------------
	查询todo: todo:get@ \n
	新建todo: todo:add@明年我要增高！ \n
	完成todo: todo:done@<todo_id> \n
	删除todo: todo:del@<todo_id> \n 
	更新todo: todo:update@<todo_id>|new_content
`

func ToDoCmd(ctx context.Context, cmd, userID, content string) string {
	var (
		meta string
	)
	switch cmd {
	case ToDoGet:
		data, err := Cli().Get(ctx, userID)
		if err != nil {
			meta = "查询 TODO 任务发生错误：" + err.Error()
		} else {
			metaStr, err := json.Marshal(data)
			if err != nil {
				meta = "序列化数据出错： " + err.Error()
			} else {
				meta = string(metaStr)
			}
		}
	case ToDoADD:
		err := Cli().Create(ctx, userID, content)
		if err != nil {
			meta = "添加TODO任务失败：" + err.Error()
		} else {
			meta = "添加TODO任务成功"
		}
	case ToDoUpdate:
		err := Cli().Update(ctx, userID, content)
		if err != nil {
			meta = "更新TODO任务失败：" + err.Error()
		} else {
			meta = "更新TODO任务成功"
		}
	case ToDoDEL:
		err := Cli().Delete(ctx, userID, content)
		if err != nil {
			meta = "删除TODO任务失败：" + err.Error()
		} else {
			meta = "删除TODO任务成功"
		}
	case ToDoDone:
		err := Cli().Done(ctx, userID, content)
		if err != nil {
			meta = "完成TODO任务失败：" + err.Error()
		} else {
			meta = "已经完成TODO任务"
		}
	default:
		meta = "命名格式有问题，请输入 HELP 查看帮助"
	}
	return meta
}
