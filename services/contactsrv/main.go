package main

import (
	"log"
	"net/http"

	_ "github.com/vnotes/workweixin/services/appsrv/dbs"
	"github.com/vnotes/workweixin/services/contactsrv/apis"

	"github.com/gorilla/mux"
)

func main() {
	var r = mux.NewRouter()

	apis.NewRouters(r)

	log.Print("listening port 11110")

	if err := http.ListenAndServe(":11110", r); err != nil {
		log.Fatalf("failed to listen server %v", err)
	}

}
