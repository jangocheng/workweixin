package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouters(r *mux.Router) {
	r.HandleFunc("/api/wx/app/reply/", WXAppAutoReply).Methods(http.MethodGet, http.MethodPost)
}
