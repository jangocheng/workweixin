package contact

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"
	"github.com/vnotes/workweixin_app/cores"
)

func WXContactAutoMated(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	log.Printf("contact manager %s", string(body))
	if err := r.ParseForm(); err != nil {
		log.Printf("parse form error %s", err)
		cores.WriteServerError(w)
		return
	}

	receiverID := cores.GetConfig().CorPID

	token := cores.GetConfig().Token
	aesKey := cores.GetConfig().AesKey

	wxCpt := wxbizmsgcrypt.NewWXBizMsgCrypt(token, aesKey, receiverID, wxbizmsgcrypt.XmlType)

	switch r.Method {
	case http.MethodGet:
		cores.WXPing(w, r, wxCpt)
	case http.MethodPost:
		WXContactManager(w, r, wxCpt)
	default:
		log.Printf("server receive http method %s which is not supported", r.Method)
		return
	}
}

func WXContactManager(w http.ResponseWriter, r *http.Request, wxCpt *wxbizmsgcrypt.WXBizMsgCrypt) {

}
