package contact

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"
	"github.com/vnotes/workweixin_app/cores"
)

func WXContactAutoMated(w http.ResponseWriter, r *http.Request) {
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

	message, err := getWXContactMsg(msg)
	if err != nil {
		log.Printf("msg %s unmarshal error %#v", string(msg), err)
		cores.WriteServerError(w)
		return
	}
	log.Printf("receive data %#v", message)

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

func getWXContactMsg(msg []byte) (*wxContactMsg, error) {
	data := &wxContactMsg{}
	if err := xml.Unmarshal(msg, data); err != nil {
		return nil, err
	}
	return data, nil
}

type wxContactMsg struct {
	ToUserName     string `xml:"ToUserName"`
	FromUserName   string `xml:"FromUserName"`
	NewUserID      string `xml:"NewUserID"`
	Department     string `xml:"Department"`
	IsLeaderInDept string `xml:"IsLeaderInDept"`
	Event          string `xml:"Event"`
	ChangeType     string `xml:"ChangeType"`

	UserID     string `xml:"UserID" db:"user_id"`
	Name       string `xml:"Name" db:"user_name"`
	Mobile     string `xml:"Mobile" db:"mobile"`
	Email      string `xml:"Email" db:"email"`
	Gender     int    `xml:"Gender" db:"gender"`
	Status     int    `xml:"Status" db:"state"`
	CreateTime int64  `xml:"CreateTime" db:"create_time"`
}
