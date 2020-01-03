package main

import (
	"log"
	"net/http"

	"github.com/vnotes/workweixin/services/appsrv/apis"
	_ "github.com/vnotes/workweixin/services/appsrv/apis/todos"
	_ "github.com/vnotes/workweixin/services/appsrv/conf"
	_ "github.com/vnotes/workweixin/services/appsrv/dbs"
	_ "github.com/vnotes/workweixin/services/appsrv/schedules"

	"github.com/gorilla/mux"
)

// todo gracefully shutdown and more
func main() {
	var r = mux.NewRouter()
	apis.NewRouters(r)

	log.Print("listening port 11111")

	if err := http.ListenAndServe(":11111", r); err != nil {
		log.Fatalf("failed to listen server %v", err)
	}
}
