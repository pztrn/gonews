package greeting

import (
	// stdlib
	"log"

	// local
	"sources.dev.pztrn.name/gonews/gonews/eventer"
	"sources.dev.pztrn.name/gonews/gonews/networker"
)

func Initialize() {
	log.Println("Initializing greeting command...")

	eventer.AddEventHandler(&eventer.EventHandler{
		Command: "internal/greeting",
		Handler: handler,
	})
}

func handler(data interface{}) interface{} {
	return &networker.Reply{Code: "201", Data: "NNTP server is ready, posting prohibited\r\n"}
}
