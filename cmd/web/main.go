package main

import (
	"log"
	"net/http"

	"github.com/rayinaw/ws-chat-basic/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("Server listening on port 5000")
	_ = http.ListenAndServe(":5000", mux)
}
