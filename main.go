package main

import (
	"log"
	"net/http"

	"github.com/vnotes/workweixin_app/cores"
)

func main() {
	r := cores.NewRouters()
	log.Fatal(http.ListenAndServe(":11111", r))
}
