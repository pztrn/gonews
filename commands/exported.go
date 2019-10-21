package commands

import (
	// stdlib
	"log"

	// local
	"sources.dev.pztrn.name/gonews/gonews/commands/capabilities"
	"sources.dev.pztrn.name/gonews/gonews/commands/greeting"
	"sources.dev.pztrn.name/gonews/gonews/commands/quit"
)

func Initialize() {
	log.Println("Initializing commands...")

	greeting.Initialize()
	quit.Initialize()
	capabilities.Initialize()
}
