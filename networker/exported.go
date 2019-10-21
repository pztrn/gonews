package networker

import (
	// stdlib
	"log"

	// local
	"sources.dev.pztrn.name/gonews/gonews/configuration"
)

// Initialize initializes package.
func Initialize() {
	log.Println("Initializing network connections handler...")

	for _, iface := range configuration.Cfg.Network {
		go startServer(iface)
	}
}
