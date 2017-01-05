package main

import (
	"log"
	"net/http"

	"github.com/dfernandez/gopi/config"
	"github.com/dfernandez/gopi/src/server"
	"github.com/dfernandez/gopi/src/commands/button_timer_on"
)

func main() {
	server := server.NewServer()

	// Register commands
	server.RegisterCommand(button_timer_on.NewButtonTimerOn())

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./public")))
	mux.Handle("/ws", server)

	log.Println("Listening on", config.SrvAddr)
	err := http.ListenAndServe(config.SrvAddr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
