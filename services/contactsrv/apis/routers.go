package apis

import (
	"net/http"

	"github.com/vnotes/workweixin/services/contactsrv/apis/users"

	"github.com/gorilla/mux"
)

func NewRouters(r *mux.Router) {
	r.HandleFunc("/api/wx/contact/", users.WXContactAutoMated).Methods(http.MethodGet, http.MethodPost)
}
