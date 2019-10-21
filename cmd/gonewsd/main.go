package main

import (
	// stdlib
	"log"
	"os"
	"os/signal"
	"syscall"

	// local
	"sources.dev.pztrn.name/gonews/gonews/commands"
	"sources.dev.pztrn.name/gonews/gonews/configuration"
	"sources.dev.pztrn.name/gonews/gonews/database"
	"sources.dev.pztrn.name/gonews/gonews/eventer"
	"sources.dev.pztrn.name/gonews/gonews/networker"
)

func main() {
	log.Println("Starting gonewsd...")

	configuration.Initialize()
	database.Initialize()
	eventer.Initialize()
	commands.Initialize()
	networker.Initialize()

	eventer.InitializeCompleted()
	log.Println("gonewsd is up and ready to serve")

	// CTRL+C handler.
	signalHandler := make(chan os.Signal, 1)
	shutdownDone := make(chan bool, 1)

	signal.Notify(signalHandler, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalHandler
		log.Println("CTRL+C or SIGTERM received, shutting down gonewsd...")
		database.Shutdown()
		shutdownDone <- true
	}()

	<-shutdownDone
	log.Println("gonewsd done")
	os.Exit(0)
}
