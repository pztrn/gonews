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

// This structure represents single NNTP client connection.
type connection struct {
	// Connection details and handlers.
	conn       net.Conn
	remoteAddr net.Addr

	// Read and write buffers
	reader *bufio.Reader
	writer *bufio.Writer
	// Scanner that helps us to read incoming data.
	readScanner *bufio.Scanner

	// Connection flags.
	// Are we in READER or MODE-READER (transit) mode?
	// Right now transit mode isn't implemented. Implementation will
	// require start using two goroutines for handling connections,
	// one for writing and one for reading.
	transit bool
	// Connection capabilites.
	capabilities []string
}

// Initialize initializes necessary things.
func (c *connection) Initialize(conn net.Conn) {
	c.conn = conn
	c.remoteAddr = c.conn.RemoteAddr()

	log.Printf("accepted connection from %v\n", conn.RemoteAddr())

	// Create buffers.
	c.reader = bufio.NewReader(conn)
	c.writer = bufio.NewWriter(conn)
	c.readScanner = bufio.NewScanner(c.reader)

	// Get capabilities for this connection.
	caps, _ := eventer.LaunchEvent("internal/capabilities", nil)
	c.capabilities = caps.([]string)

	// Set transit mode by default, according to RFC.
	c.transit = true
}

// Start starts working with connection. Should be launched in separate
// goroutine.
// It will send greeting and then falls into infinite loop for working
// with connection until the end.
// Right now it implements only READER mode, no transit (which is used
// by server-to-server peering extensively).
func (c *connection) Start() {
	defer func() {
		err := c.conn.Close()
		if err != nil {
			log.Println("Failed to close connection from " + c.remoteAddr.String() + ": " + err.Error())
		}
		log.Println("Connection from " + c.remoteAddr.String() + " closed")
	}()

	// Send greeting.
	greetingData, _ := eventer.LaunchEvent("internal/greeting", nil)
	greetingReply := greetingData.(*Reply)

	_, err := c.writer.WriteString(greetingReply.Code + " " + greetingReply.Data)
	if err != nil {
		log.Println("Failed to write greeting for " + c.remoteAddr.String() + ": " + err.Error())
		return
	}
	c.writer.Flush()

	// Start reading for commands.
	// Every command can be represented as slice where first element
	// is actual command and all next - parameters.
	// By default we read only one line per iteration.
	// ToDo: multiline data parser for posting.
	for {
		dataAppeared := c.readScanner.Scan()
		if !dataAppeared {
			log.Println("Failed to read data from " + c.remoteAddr.String() + ": " + c.readScanner.Err().Error())
			break
		}

		log.Println("Got data: " + c.readScanner.Text())

		// ToDo: what if we'll upload binary data here?
		// Not supported yet.
		data := strings.Split(c.readScanner.Text(), " ")

		// Separate capabilities worker.
		if strings.ToLower(data[0]) == "capabilities" {
			dataToWrite := "Capability list:\r\n"

			for _, cap := range c.capabilities {
				dataToWrite += cap + "\r\n"
			}

			// We're also mode-switching server (in future), so we should
			// also be aware of mode reader things. Writing to client will
			// depend on c.transit variable.
			// We will announce MODE-READER capability after initial
			// connection (because we're in transit mode by default, according
			// to RFC), and when client issue "MODE READER" command we will
			// stop announcing MODE-READER and will start announce READER
			// capability.
			if c.transit {
				dataToWrite += "MODE-READER\r\n"
			} else {
				dataToWrite += "READER\r\n"
			}

			dataToWrite += ".\r\n"
			c.writer.WriteString(dataToWrite)
			c.writer.Flush()
			continue
		}

		// Mode worker. Reader only now.
		if strings.ToLower(data[0]) == "mode" && strings.ToLower(data[1]) == "reader" {
			c.transit = false
			// In any case we'll require user authentication for posting.
			c.writer.WriteString("201 Posting prohibited\r\n")
			c.writer.Flush()
			continue
		}

		// Execute passed command.
		replyRaw, err := eventer.LaunchEvent("commands/"+strings.ToLower(data[0]), data[1:])
		if err != nil {
			// We won't break here as this is just logging of appeared error.
			log.Println("Error appeared while processing command '" + data[0] + "' for " + c.remoteAddr.String() + ": " + err.Error())
		}

		// We might have nil in reply, so we'll assume that passed command
		// is unknown to us.
		if replyRaw == nil {
			_, err := c.writer.WriteString(unknownCommandErrorCode + " " + unknownCommandErrorText + "\r\n")
			if err != nil {
				log.Println("Failed to write string to socket for " + c.remoteAddr.String() + ": " + err.Error())
				break
			}
			c.writer.Flush()
			continue
		}

		// Every reply will be a reply struct.
		reply := replyRaw.(*Reply)

		_, err1 := c.writer.WriteString(reply.Code + " " + reply.Data)
		if err1 != nil {
			log.Println("Failed to write string to socket for " + c.remoteAddr.String() + ": " + err1.Error())
			break
		}
		c.writer.Flush()

		// Check for QUIT command.
		if strings.ToLower(data[0]) == "quit" {
			log.Println("QUIT command received, closing connection to " + c.remoteAddr.String())
			break
		}
	}
}
