package main

import (
	"log"
	"net/http"

	"automated-bot/cores"
)

func main() {
	r := cores.NewRouters()
	log.Fatal(http.ListenAndServe(":11111", r))
}
