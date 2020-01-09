package users

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vnotes/workweixin/services/contactsrv/conf"
	"github.com/vnotes/workweixin/services/cores"

	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"
)

func getContactRequestParameter(r *http.Request) *cores.WXPing {
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

func WXContactAutoMated(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("parse form error %s", err)
		cores.WriteServerError(w)
		return
	}

	receiverID := conf.Conf.CorPID

	token := conf.Conf.Token
	aesKey := conf.Conf.AesKey

	wxCpt := wxbizmsgcrypt.NewWXBizMsgCrypt(token, aesKey, receiverID, wxbizmsgcrypt.XmlType)

	reqParam := getContactRequestParameter(r)

	switch r.Method {
	case http.MethodGet:
		rsp, err := cores.WXPong(reqParam, wxCpt)
		if err != nil {
			cores.WriteServerError(w)
			return
		}
		cores.WriteServerSuccess(w, rsp)
	case http.MethodPost:
		WXContactManager(w, r, wxCpt)
	default:
		log.Printf("server receive http method %s which is not supported", r.Method)
		return
	}
}

func WXContactManager(w http.ResponseWriter, r *http.Request, wx *wxbizmsgcrypt.WXBizMsgCrypt) {
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
	log.Printf("receive data %s", string(msg))

	message, err := getWXContactMsg(msg)
	if err != nil {
		log.Printf("msg %s unmarshal error %#v", string(msg), err)
		cores.WriteServerError(w)
		return
	}

	var eventErr error
	switch message.ChangeType {
	case CreateUser:
		eventErr = Cli().CreateUser(r.Context(), message)
	case UpdateUser:
		eventErr = Cli().UpdateUser(r.Context(), message)
	case DeleteUser:
		eventErr = Cli().DeleteUser(r.Context(), message)
	default:
		return
	}
	//nolint:staticcheck
	if eventErr != nil {
	}
}

func getWXContactMsg(msg []byte) (*WXContactMsg, error) {
	data := &WXContactMsg{}
	if err := xml.Unmarshal(msg, data); err != nil {
		return nil, err
	}
	return data, nil
}

type WXContactMsg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	Event        string `xml:"Event"`
	ChangeType   string `xml:"ChangeType"`

	UserID     string  `xml:"UserID"`
	NewUserID  *string `xml:"NewUserID"`
	Name       *string `xml:"Name"`
	Mobile     *string `xml:"Mobile"`
	Email      *string `xml:"Email"`
	Gender     *int    `xml:"Gender"`
	Status     *int    `xml:"Status"`
	CreateTime int64   `xml:"CreateTime"`
}
