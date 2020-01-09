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

func WriteServerSuccess(w http.ResponseWriter, rsp []byte)  {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(rsp)
}

type WXPing struct {
	MsgSignature string
	TimeStamp    string
	Nonce        string
	Echo         string
}

func WXPong(w http.ResponseWriter, ping *WXPing, wx *wxbizmsgcrypt.WXBizMsgCrypt) {
	echoStr, cryptErr := wx.VerifyURL(ping.MsgSignature, ping.TimeStamp, ping.Echo, ping.Echo)
	if cryptErr != nil {
		log.Printf("verify error %+v", cryptErr)
		WriteServerError(w)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(echoStr)
}

type Response struct {
	ErrCode int
	ErrMsg  string
}
