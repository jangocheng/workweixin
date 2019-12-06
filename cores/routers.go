package cores

import (
	"net/http"

	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func NewRouters() *mux.Router {
	return r
}

func init() {
	r.HandleFunc("/api/wx/ping/", WXPing).Methods(http.MethodGet)
}
