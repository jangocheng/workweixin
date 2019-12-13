package main

import (
	"log"
	"net/http"

	"github.com/vnotes/workweixin_app/cores/app"
	"github.com/vnotes/workweixin_app/cores/contact"

	"github.com/gorilla/mux"
)

func main() {

	var r = mux.NewRouter()

	app.NewRouters(r)
	contact.NewRouters(r)
	log.Fatal(http.ListenAndServe(":11111", r))
}
