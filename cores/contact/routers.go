package contact

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouters(r *mux.Router) {
	r.HandleFunc("/api/wx/contact/", WXContactAutoMated).Methods(http.MethodGet, http.MethodPost)
}
