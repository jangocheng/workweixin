package cores

import (
	"log"
	"net/http"

	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"
)

func WriteServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("server error"))
}

func WXPong(w http.ResponseWriter, r *http.Request, wx *wxbizmsgcrypt.WXBizMsgCrypt) {
	sig := r.Form.Get("msg_signature")
	timeStamp := r.Form.Get("timestamp")
	nonce := r.Form.Get("nonce")
	echo := r.Form.Get("echostr")

	echoStr, cryptErr := wx.VerifyURL(sig, timeStamp, nonce, echo)
	if cryptErr != nil {
		log.Printf("verify error %+v", cryptErr)
		WriteServerError(w)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(echoStr)
}

type AccessToken struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}