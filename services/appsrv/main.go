package main

import (
	"log"
	"net/http"

	"github.com/vnotes/workweixin/services/appsrv/apis"
	"github.com/vnotes/workweixin/services/appsrv/apis/todos"
	"github.com/vnotes/workweixin/services/appsrv/conf"
	"github.com/vnotes/workweixin/services/appsrv/dbs"
	_ "github.com/vnotes/workweixin/services/appsrv/schedules"

	"github.com/gorilla/mux"
)

func main() {
	conf.InitConfig()
	todos.InitToDoGRPC()
	dbs.InitMySQL()
	dbs.InitRedis()

	var r = mux.NewRouter()
	apis.NewRouters(r)

	log.Print("listening port 11111")

	if err := http.ListenAndServe(":11111", r); err != nil {
		log.Fatalf("failed to listen server %v", err)
	}
}
