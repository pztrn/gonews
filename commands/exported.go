package commands

import (
	// stdlib
	"log"

	// local
	"develop.pztrn.name/gonews/gonews/commands/capabilities"
	"develop.pztrn.name/gonews/gonews/commands/greeting"
	"develop.pztrn.name/gonews/gonews/commands/quit"
)

func Initialize() {
	log.Println("Initializing commands...")

	greeting.Initialize()
	quit.Initialize()
	capabilities.Initialize()
}
