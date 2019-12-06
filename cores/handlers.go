package cores

import (
	"log"
	"net/http"
	"os"

	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"
)

func WXPing(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("parse form error %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(""))
		return
	}

	msgSig := r.Form.Get("msg_signature")
	timeStamp := r.Form.Get("timestamp")
	echo := r.Form.Get("echostr")
	nonce := r.Form.Get("nonce")

	token := os.Getenv("TOKEN")
	receiverID := os.Getenv("RECEIVER_ID")
	aesKey := os.Getenv("AES_KEY")

	wxCpt := wxbizmsgcrypt.NewWXBizMsgCrypt(token, aesKey, receiverID, wxbizmsgcrypt.XmlType)
	echoStr, cryptErr := wxCpt.VerifyURL(msgSig, timeStamp, nonce, echo)
	if cryptErr != nil {
		log.Printf("verify error %+v", cryptErr)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("server error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(echoStr)
}
