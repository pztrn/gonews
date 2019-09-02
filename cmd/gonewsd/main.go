package main

import (
	// stdlib
	"log"
	"os"
	"os/signal"
	"syscall"

	// local
	"develop.pztrn.name/gonews/gonews/eventer"
	"develop.pztrn.name/gonews/gonews/networker"
)

func main() {
	log.Println("Starting gonewsd...")

	eventer.Initialize()
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
		shutdownDone <- true
	}()

	<-shutdownDone
	log.Println("gonewsd done")
	os.Exit(0)
}
