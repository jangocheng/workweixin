package app

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/vnotes/workweixin_app/cores"
	"github.com/vnotes/workweixin_app/cores/todos"

	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"
)

func WXAppAutoReply(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("parse form error %s", err)
		cores.WriteServerError(w)
		return
	}

	receiverID := Conf.CorPID

	token := Conf.Token
	aesKey := Conf.AesKey

	wxCpt := wxbizmsgcrypt.NewWXBizMsgCrypt(token, aesKey, receiverID, wxbizmsgcrypt.XmlType)

	switch r.Method {
	case http.MethodGet:
		cores.WXPong(w, r, wxCpt)
	case http.MethodPost:
		wxAutoReplyMsg(w, r, wxCpt)
	default:
		log.Printf("server receive http method %s which is not supported", r.Method)
		return
	}
}

func wxAutoReplyMsg(w http.ResponseWriter, r *http.Request, wx *wxbizmsgcrypt.WXBizMsgCrypt) {
	sig := r.Form.Get("msg_signature")
	timeStamp := r.Form.Get("timestamp")
	nonce := r.Form.Get("nonce")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("server read body error %#v", err)
		cores.WriteServerError(w)
		return
	}
	msg, cryptErr := wx.DecryptMsg(sig, timeStamp, nonce, body)
	if cryptErr != nil {
		log.Printf("decode error %#v", cryptErr)
		cores.WriteServerError(w)
		return
	}
	log.Printf("app receive data %s", string(msg))

	message, err := getWXAppMsg(msg)
	if err != nil {
		log.Printf("msg %s unmarshal error %#v", string(msg), err)
		cores.WriteServerError(w)
		return
	}

	var rspMsg string

	rawMsg := message.Content
	if rawMsg == "HELP" {
		rspMsg = todos.HELP
	} else {
		if strings.HasPrefix(rawMsg, "todo:") && strings.Contains(rawMsg, "@") {
			var (
				content = strings.Split(message.Content, "@")
				cmd     = content[0]
				userID  = message.FromUserName
				text    = content[1]
			)
			if true {
				rspMsg = todos.ToDoCmdRPC(r.Context(), cmd, userID, text)
			} else {
				rspMsg = todos.ToDoCmd(r.Context(), cmd, userID, text)
			}
		}
	}

	if rspMsg == "" {
		return
	}

	replyMsgRsp := &WXAppMsg{
		ToUserName:   message.ToUserName,
		FromUserName: message.FromUserName,
		CreateTime:   message.CreateTime,
		MsgType:      message.MsgType,
		Content:      rspMsg,
		MsgId:        message.MsgId,
		AgentID:      message.AgentID,
	}
	msgByte, err := xml.Marshal(replyMsgRsp)
	if err != nil {
		log.Printf("marshal data %s error %#v", string(msgByte), err)
		cores.WriteServerError(w)
		return
	}
	rsp, cryptErr := wx.EncryptMsg(string(msgByte), timeStamp, nonce)
	if cryptErr != nil {
		log.Printf("encode data %s error %#v", string(msg), cryptErr)
		cores.WriteServerError(w)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(rsp)
}

func getWXAppMsg(msg []byte) (*WXAppMsg, error) {
	data := &WXAppMsg{}
	if err := xml.Unmarshal(msg, data); err != nil {
		return nil, err
	}
	return data, nil
}

type WXAppMsg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        string `xml:"MsgId"`
	AgentID      string `xml:"AgentID"`
}
