package main

import (
	"log"
	"net/http"

	"github.com/vnotes/workweixin/services/contactsrv/apis"
	"github.com/vnotes/workweixin/services/contactsrv/conf"
	"github.com/vnotes/workweixin/services/contactsrv/dbs"
	"github.com/vnotes/workweixin/services/contactsrv/tracings"

	"github.com/gorilla/mux"
)

const ServiceName = "weixin.contactsrv"

func main() {
	conf.InitConfig()
	dbs.InitMySQL()

	tracings.InitTracing(ServiceName)

	defer func() {
		_ = tracings.CloseTracer()
	}()
	var r = mux.NewRouter()

	apis.NewRouters(r)

	log.Print("listening port 11110")

	if err := http.ListenAndServe(":11110", r); err != nil {
		log.Fatalf("failed to listen server %v", err)
	}

}
