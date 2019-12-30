package main

import (
	"log"
	"net/http"

	"github.com/vnotes/workweixin_app/cores/app"
	"github.com/vnotes/workweixin_app/cores/contact"
	_ "github.com/vnotes/workweixin_app/cores/dbs"
	"github.com/vnotes/workweixin_app/cores/schedules"
	"github.com/vnotes/workweixin_app/cores/todos"

	"github.com/gorilla/mux"
)

// todo gracefully shutdown and more
func main() {
	InitAction()

	defer func() {
		DeferAction()
	}()
}

func DeferAction() {
	todos.CloseGRPCClient()
}

func InitAction() {

	schedules.RegisterCronJob()

	var r = mux.NewRouter()
	app.NewRouters(r)
	contact.NewRouters(r)

	go todos.InitGRPCServer()
	todos.InitGRPCClient()

	log.Print("listening port 11111")
	if err := http.ListenAndServe(":11111", r); err != nil {
		log.Fatalf("failed to listen server %v", err)
	}

}
