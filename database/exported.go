package database

import (
	// stdlib
	"log"
	"time"

	// other
	"github.com/jmoiron/sqlx"
)

var (
	// Conn is database connection.
	Conn *sqlx.DB

	// Shutdown flags.
	// Sets to true when Shutdown() is called to indicate other subsystes
	// that we're shutting down.
	weAreShuttingDown bool
	// Sets to true when connection watcher will be stopped.
	connWatcherStopped bool
)

func Initialize() {
	log.Println("Initializing database handler...")

	// Reset variables to their default startup state because they
	// can be set to other values while executing tests.
	Conn = nil
	weAreShuttingDown = false
	connWatcherStopped = false

	go startConnectionWatcher()
}

func Shutdown() {
	weAreShuttingDown = true
	for {
		if connWatcherStopped {
			break
		}
	}

	time.Sleep(time.Millisecond * 500)
}
