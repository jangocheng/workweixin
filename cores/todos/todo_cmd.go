package todos

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
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

func ToDoCmd(ctx context.Context, cmd, userID, content string) (meta string) {
	switch cmd {
	case ToDoGet:
		data, err := Cli().Get(ctx, userID)
		if err != nil {
			meta = "查询 TODO 任务发生错误：" + err.Error()
			return
		}
		metaStr, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			meta = "序列化数据出错： " + err.Error()
			return
		}
		meta = "查询 TODO LIST 成功\n" + string(metaStr)
		return
	case ToDoADD:
		err := Cli().Create(ctx, userID, content)
		if err != nil {
			meta = "添加TODO任务失败：" + err.Error()
			return
		}
		meta = "添加TODO任务成功"
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
		updateContent := text[1]
		err = Cli().Update(ctx, userID, updateContent, todoID)
		if err != nil {
			meta = "更新TODO任务失败：" + err.Error()
			return
		}
		meta = "更新TODO任务成功"
		return
	case ToDoDEL:
		todoID, err := strconv.ParseUint(content, 10, 64)
		if err != nil {
			meta = "命令有误，请输入 HELP"
			return
		}
		err = Cli().Delete(ctx, userID, todoID)
		if err != nil {
			meta = "删除TODO任务失败：" + err.Error()
			return
		}
		meta = "删除TODO任务成功"
		return
	case ToDoDone:
		todoID, err := strconv.ParseUint(content, 10, 64)
		if err != nil {
			meta = "命令有误，请输入 HELP"
			return
		}
		err = Cli().Done(ctx, userID, todoID)
		if err != nil {
			meta = "完成TODO任务失败：" + err.Error()
			return
		}
		meta = "已经完成TODO任务"
		return
	default:
		meta = "命名格式有问题，请输入 HELP 查看帮助"
	}
	return meta
}
