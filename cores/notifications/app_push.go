package notifications

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/vnotes/workweixin_app/cores"
	"github.com/vnotes/workweixin_app/cores/app"
)

type Notifier interface {
	AppMsgPush()
}

type appMsgReq struct {
	ToUser  string     `json:"touser"`
	MsgType string     `json:"msgtype"`
	AgentID int64      `json:"agentid"`
	Text    MsgContent `json:"text"`
}
type MsgContent struct {
	Content string `json:"content"`
}

func AppMsgPush() {
	appConf := app.Conf
	token := cores.GetAccessToken(appConf.CorPID, appConf.Secret)
	if token == "" {
		return
	}
	agentID, err := strconv.ParseInt(appConf.AgentID, 10, 64)
	if err != nil {
		return
	}

	uri := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + token
	pushData := &appMsgReq{
		ToUser:  "zaizai",
		MsgType: "text",
		AgentID: agentID,
		Text:    MsgContent{Content: "text"},
	}
	body, err := json.Marshal(pushData)
	if err != nil {
		return
	}
	meta := bytes.NewBuffer(body)
	rsp := &cores.Response{}
	err = cores.InitClient("POST", uri, meta).HttpDo(rsp)
	if err != nil {
		return
	}
	if rsp.ErrMsg != "ok" {
		return
	}
}
