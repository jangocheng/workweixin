package main

import (
	"github.com/vnotes/workweixin_app/cores/contact"
	"log"
	"net/http"
	"os"

	"github.com/vnotes/workweixin_app/cores"
	"github.com/vnotes/workweixin_app/cores/application"

	"github.com/gorilla/mux"
)

func main() {
	checkEnvVars()

	var r = mux.NewRouter()

	application.NewRouters(r)
	contact.NewRouters(r)
	log.Fatal(http.ListenAndServe(":11111", r))
}

func checkEnvVars() {
	corPID := os.Getenv("Cor_PID")

	corSecret := os.Getenv("Cor_Secret")
	agentID := os.Getenv("AGENT_ID")
	token := os.Getenv("TOKEN")
	aesKey := os.Getenv("AES_KEY")

	if corPID == "" || corSecret == "" || agentID == "" || token == "" || aesKey == "" {
		panic("init env vars failed")
	}
	cores.SetConfig(corPID, corSecret, agentID, token, aesKey)
}
