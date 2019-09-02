package networker

import (
	// stdlib
	"log"
	"net"

	// local
	"develop.pztrn.name/gonews/gonews/configuration"
)

// This function responsible for accepting incoming connections for
// each address configuration.
func startServer(config configuration.Network) {
	log.Println("Starting server on " + config.Address + " (type: " + config.Type + ")")

	l, err := net.Listen("tcp", config.Address)
	if err != nil {
		log.Fatalln("Failed to start TCP server on " + config.Address + ": " + err.Error())
	}
	defer func() {
		err := l.Close()
		if err != nil {
			log.Println("Failed to close TCP server on " + config.Address + ": " + err.Error())
		}
	}()

	for {
		conn, err1 := l.Accept()
		if err1 != nil {
			log.Println("Failed to accept new incoming connection: " + err1.Error())
			continue
		}

		go connectionWorker(conn)
	}
}
