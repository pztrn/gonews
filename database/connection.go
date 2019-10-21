package database

import (
	// stdlib
	"log"
	"strings"
	"time"

	// local
	"develop.pztrn.name/gonews/gonews/configuration"
	"develop.pztrn.name/gonews/gonews/database/migrations"

	// other
	"github.com/jmoiron/sqlx"
	// PostgreSQL driver.
	_ "github.com/lib/pq"
)

func startConnectionWatcher() {
	log.Println("Initializing database connection watcher...")

	migrations.Initialize()

	// Checking configuration validity.
	// Parameters should not be specified in DSN.
	if strings.Contains(configuration.Cfg.Database.DSN, "?") {
		log.Fatalln("Database DSN should not contain parameters, specify them in DATABASE_PARAMS environment variable!")
	}

	// DSN should be defined.
	if configuration.Cfg.Database.DSN == "" {
		log.Fatalln("Database DSN should be defined!")
	}

	ticker := time.NewTicker(time.Second * time.Duration(configuration.Cfg.Database.Timeout))

	// First start - manually.
	_ = watcher()

	// Then - every ticker tick.
	for range ticker.C {
		doBreak := watcher()
		if doBreak {
			break
		}
	}

	ticker.Stop()
	log.Println("Connection watcher stopped and connection to database was shutted down")

	connWatcherStopped = true
}

// Actual connection watcher. Returns true if we should stop watching
// (e.g. due to shutdown) and false if everything is okay.
func watcher() bool {
	// If we're shutting down - stop connection watcher.
	if weAreShuttingDown {
		log.Println("Closing database connection...")

		err := Conn.Close()
		if err != nil {
			log.Println("Failed to close database connection")
		}

		Conn = nil

		return true
	}

	// If connection is nil - try to establish (or reestablish)
	// connection.
	if Conn == nil {
		log.Println("(Re)Establishing connection to PostgreSQL...")

		// Compose DSN.
		dsn := configuration.Cfg.Database.DSN
		if configuration.Cfg.Database.Params != "" {
			dsn += "?" + configuration.Cfg.Database.Params
		}

		// Connect to database.
		dbConn, err := sqlx.Connect("postgres", dsn)
		if err != nil {
			log.Fatalf("Failed to connect to PostgreSQL database, will try to reconnect after %d seconds\n", configuration.Cfg.Database.Timeout)
			return false
		}

		log.Println("Database connection (re)established")

		// Migrate database.
		migrations.Migrate(dbConn.DB)

		Conn = dbConn
	}

	return false
}
