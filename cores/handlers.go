package cores

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"
)

func WXAutoReply(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("parse form error %s", err)
		writeServerError(w)
		return
	}

	token := os.Getenv("TOKEN")
	receiverID := os.Getenv("RECEIVER_ID")
	aesKey := os.Getenv("AES_KEY")

	wxCpt := wxbizmsgcrypt.NewWXBizMsgCrypt(token, aesKey, receiverID, wxbizmsgcrypt.XmlType)

	switch r.Method {
	case http.MethodGet:
		wxPing(w, r, wxCpt)
	case http.MethodPost:
		wxAutoReplyMsg(w, r, wxCpt)
	default:
		log.Printf("server receive http method %s which is not supported", r.Method)
		return
	}
}

func wxPing(w http.ResponseWriter, r *http.Request, wx *wxbizmsgcrypt.WXBizMsgCrypt) {
	sig := r.Form.Get("msg_signature")
	timeStamp := r.Form.Get("timestamp")
	nonce := r.Form.Get("nonce")
	echo := r.Form.Get("echostr")

	echoStr, cryptErr := wx.VerifyURL(sig, timeStamp, nonce, echo)
	if cryptErr != nil {
		log.Printf("verify error %+v", cryptErr)
		writeServerError(w)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(echoStr)
}

func wxAutoReplyMsg(w http.ResponseWriter, r *http.Request, wx *wxbizmsgcrypt.WXBizMsgCrypt) {
	sig := r.Form.Get("msg_signature")
	timeStamp := r.Form.Get("timestamp")
	nonce := r.Form.Get("nonce")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("server read body error %#v", err)
		writeServerError(w)
		return
	}
	msg, cryptErr := wx.DecryptMsg(sig, timeStamp, nonce, body)
	if cryptErr != nil {
		log.Printf("decode error %#v", cryptErr)
		writeServerError(w)
		return
	}
	message, err := getWXAppMsg(msg)
	if err != nil {
		log.Printf("msg %s unmarshal error %#v", string(msg), err)
		writeServerError(w)
		return
	}
	rspMsg := "auto-reply-source-message:\n" + message.Content
	replyMsgRsp := &wxAppMsg{
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
		writeServerError(w)
		return
	}
	rsp, cryptErr := wx.EncryptMsg(string(msgByte), timeStamp, nonce)
	if cryptErr != nil {
		log.Printf("encode data %s error %#v", string(msg), cryptErr)
		writeServerError(w)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(rsp)
}

func getWXAppMsg(msg []byte) (*wxAppMsg, error) {
	data := &wxAppMsg{}
	if err := xml.Unmarshal(msg, data); err != nil {
		return nil, err
	}
	return data, nil
}
