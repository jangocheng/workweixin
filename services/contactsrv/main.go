package main

import (
	"log"
	"net/http"

	"github.com/vnotes/workweixin/services/contactsrv/apis"
	"github.com/vnotes/workweixin/services/contactsrv/conf"
	"github.com/vnotes/workweixin/services/contactsrv/dbs"

	"github.com/gorilla/mux"
)

func main() {
	conf.InitConfig()
	dbs.InitMySQL()

	var r = mux.NewRouter()

	apis.NewRouters(r)

	log.Print("listening port 11110")

	if err := http.ListenAndServe(":11110", r); err != nil {
		log.Fatalf("failed to listen server %v", err)
	}

}
