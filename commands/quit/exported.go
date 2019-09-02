package quit

import (
	// stdlib
	"log"

	// local
	"develop.pztrn.name/gonews/gonews/eventer"
	"develop.pztrn.name/gonews/gonews/networker"
)

func Initialize() {
	log.Println("Initializing quit command...")

	eventer.AddEventHandler(&eventer.EventHandler{
		Command: "commands/quit",
		Handler: handler,
	})
}

func handler(data interface{}) interface{} {
	return &networker.Reply{Code: "205", Data: "NNTP Service exits normally\r\n"}
}
