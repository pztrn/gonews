package capabilities

import (
	// stdlib
	"log"

	// local
	"sources.dev.pztrn.name/gonews/gonews/eventer"
)

var capabilities = []string{
	"VERSION 2",
}

func Initialize() {
	log.Println("Initializing capabilities command...")

	// Global capabilities adder.
	eventer.AddEventHandler(&eventer.EventHandler{
		Command: "internal/capability_add",
		Handler: addCapability,
	})

	eventer.AddEventHandler(&eventer.EventHandler{
		Command: "internal/capabilities",
		Handler: handler,
	})
}

func addCapability(data interface{}) interface{} {
	capabilities = append(capabilities, data.(string))

	return nil
}

//func handler(data interface{}) interface{} {
//	dataToReturn := "Capability list:\r\n"
//
//	for _, cap := range capabilities {
//		dataToReturn += cap + "\r\n"
//	}
//	dataToReturn += ".\r\n"
//	return &networker.Reply{Code: "101", Data: dataToReturn}
//}

func handler(data interface{}) interface{} {
	caps := make([]string, len(capabilities))
	copy(caps, capabilities)

	return caps
}
