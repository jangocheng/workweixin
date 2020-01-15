package apis

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/vnotes/workweixin/services/appsrv/apis/todos"
	"github.com/vnotes/workweixin/services/appsrv/conf"
	"github.com/vnotes/workweixin/services/appsrv/tracings"
	"github.com/vnotes/workweixin/services/cores"
	"github.com/vnotes/workweixin/utils"

	"github.com/opentracing/opentracing-go"
	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"
)

func getWeiXinCrypto() *wxbizmsgcrypt.WXBizMsgCrypt {
	receiverID := conf.Conf.CorPID

	token := conf.Conf.Token
	aesKey := conf.Conf.AesKey

	wxCpt := wxbizmsgcrypt.NewWXBizMsgCrypt(token, aesKey, receiverID, wxbizmsgcrypt.XmlType)
	return wxCpt
}

func getAppRequestParameter(r *http.Request) *cores.WXPing {
	sig := r.Form.Get("msg_signature")
	timeStamp := r.Form.Get("timestamp")
	nonce := r.Form.Get("nonce")
	echo := r.Form.Get("echostr")

	rsp := &cores.WXPing{
		MsgSignature: sig,
		TimeStamp:    timeStamp,
		Nonce:        nonce,
		Echo:         echo,
	}
	return rsp
}

func WXAppAutoReply(w http.ResponseWriter, r *http.Request) {
	span := tracings.Tracer.StartSpan("app-message")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(r.Context(), span)

	if err := r.ParseForm(); err != nil {
		log.Printf("parse form error %s", err)
		cores.WriteServerError(w)
		return
	}

	wxCpt := getWeiXinCrypto()

	reqParam := getAppRequestParameter(r)

	switch r.Method {
	case http.MethodGet:
		rsp, err := cores.WXPong(reqParam, wxCpt)
		if err != nil {
			cores.WriteServerError(w)
			return
		}
		cores.WriteServerSuccess(w, rsp)
	case http.MethodPost:
		message, err := decodeWeiXinMsg(r, wxCpt, reqParam)
		if err != nil {
			cores.WriteServerError(w)
			return
		}
		newMessage := wxAutoReplyMsg(ctx, message)
		response, err := encodeWeiXinMsg(newMessage, wxCpt, reqParam)
		if err != nil {
			cores.WriteServerError(w)
			return
		}
		cores.WriteServerSuccess(w, response)
	default:
		log.Printf("server receive http method %s which is not supported", r.Method)
	}
}

func decodeWeiXinMsg(r *http.Request, wx *wxbizmsgcrypt.WXBizMsgCrypt, ping *cores.WXPing) (*WXAppMsg, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("server read body error %#v", err)
		return nil, err
	}
	var (
		msg      []byte
		cryptErr *wxbizmsgcrypt.CryptError
	)
	if conf.Conf.ISDebug {
		msg = body
	} else {
		msg, cryptErr = wx.DecryptMsg(ping.MsgSignature, ping.TimeStamp, ping.Nonce, body)
		if cryptErr != nil {
			log.Printf("decode error %#v", cryptErr)
			return nil, errors.New("decode message error")
		}
	}

	log.Printf("app receive data %s", string(msg))

	message, err := getWXAppMsg(msg)
	if err != nil {
		log.Printf("msg %s unmarshal error %#v", string(msg), err)
		return nil, err
	}
	return message, nil
}

func encodeWeiXinMsg(msg *WXAppMsg, wx *wxbizmsgcrypt.WXBizMsgCrypt, ping *cores.WXPing) ([]byte, error) {
	msgByte, err := xml.Marshal(msg)
	if err != nil {
		log.Printf("marshal data %s error %#v", string(msgByte), err)
		return nil, err
	}
	var (
		rsp      []byte
		cryptErr *wxbizmsgcrypt.CryptError
	)
	if conf.Conf.ISDebug {
		rsp = msgByte
	} else {
		rsp, cryptErr = wx.EncryptMsg(string(msgByte), ping.TimeStamp, ping.Nonce)
		if cryptErr != nil {
			log.Printf("encode data %s error %#v", string(msgByte), cryptErr)
			return nil, errors.New("encode message error")
		}
	}
	return rsp, nil
}

func wxAutoReplyMsg(ctx context.Context, message *WXAppMsg) *WXAppMsg {
	var rspMsg string

	rawMsg := utils.ReplaceString(message.Content, []string{" ", "\n"})

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
			span, _ := opentracing.StartSpanFromContext(ctx, "todo-list-data")
			defer span.Finish()

			contactURL := fmt.Sprintf("http://%s:11110/api/wx/contact/pong", conf.Conf.ContactWork)
			client := cores.InitClient("GET", contactURL, nil)
			reqHeader := client.GetRequestHeader()
			if reqHeader != nil {
				_ = span.Tracer().Inject(
					span.Context(),
					opentracing.HTTPHeaders,
					opentracing.HTTPHeadersCarrier(reqHeader),
				)
				client.JustDo()
			}
			rspMsg = todos.ToDoCmd(ctx, cmd, userID, text)
		}
	}

	if rspMsg == "" {
		return nil
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
	return replyMsgRsp
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
