package networker

import (
	// stdlib
	"bufio"
	"log"
	"net"
	"strings"

	// local
	"develop.pztrn.name/gonews/gonews/eventer"
)

// This function is a connection worker.
func connectionWorker(conn net.Conn) {
	remoteAddr := conn.RemoteAddr()
	log.Printf("accepted connection from %v\n", conn.RemoteAddr())

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println("Failed to close connection from " + remoteAddr.String() + ": " + err.Error())
		}
		log.Println("Connection from " + remoteAddr.String() + " closed")
	}()

	// Create buffers.
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	scanr := bufio.NewScanner(r)

	// Send greeting.
	greetingData, _ := eventer.LaunchEvent("internal/greeting", nil)
	greetingReply := greetingData.(*Reply)

	_, err := w.WriteString(greetingReply.Code + " " + greetingReply.Data)
	if err != nil {
		log.Println("Failed to write greeting for " + remoteAddr.String() + ": " + err.Error())
		return
	}
	w.Flush()

	// Start reading for commands.
	// Every command can be represented as slice where first element
	// is actual command and all next - parameters.
	// By default we read only one line per iteration.
	// ToDo: multiline data parser for posting.
	for {
		dataAppeared := scanr.Scan()
		if !dataAppeared {
			log.Println("Failed to read data from " + remoteAddr.String() + ": " + scanr.Err().Error())
			break
		}

		log.Println("Got data: " + scanr.Text())

		// ToDo: what if we'll upload binary data here?
		// Not supported yet.
		data := strings.Split(scanr.Text(), " ")
		replyRaw, err := eventer.LaunchEvent("commands/"+strings.ToLower(data[0]), data[1:])
		if err != nil {
			// We won't break here as this is just logging of appeared error.
			log.Println("Error appeared while processing command '" + data[0] + "' for " + remoteAddr.String() + ": " + err.Error())
		}

		// We might have nil in reply, so we'll assume that passed command
		// is unknown to us.
		if replyRaw == nil {
			_, err := w.WriteString(unknownCommandErrorCode + " " + unknownCommandErrorText + "\r\n")
			if err != nil {
				log.Println("Failed to write string to socket for " + remoteAddr.String() + ": " + err.Error())
				break
			}
			w.Flush()
			continue
		}

		// Every reply will be a reply struct.
		reply := replyRaw.(*Reply)

		_, err1 := w.WriteString(reply.Code + " " + reply.Data)
		if err1 != nil {
			log.Println("Failed to write string to socket for " + remoteAddr.String() + ": " + err1.Error())
			break
		}
		w.Flush()

		// Check for QUIT command.
		if strings.ToLower(data[0]) == "quit" {
			log.Println("QUIT command received, closing connection to " + remoteAddr.String())
			break
		}
	}
}
