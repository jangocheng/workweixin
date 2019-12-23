package main

import (
	"log"
	"net/http"

	"github.com/robfig/cron/v3"
	"github.com/vnotes/workweixin_app/cores/app"
	"github.com/vnotes/workweixin_app/cores/contact"
	_ "github.com/vnotes/workweixin_app/cores/dbs"
	"github.com/vnotes/workweixin_app/cores/schedules"

	"github.com/gorilla/mux"
)

func main() {
	var c = cron.New()
	schedules.RegisterCronJob(c)

	var r = mux.NewRouter()
	app.NewRouters(r)
	contact.NewRouters(r)
	log.Fatal(http.ListenAndServe(":11111", r))
}
