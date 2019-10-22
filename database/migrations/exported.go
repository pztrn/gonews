package migrations

import (
	// stdlib
	"database/sql"
	"log"
	"os"
	"strconv"

	// other
	"github.com/pressly/goose"
)

// Initialize initializes database migrations. This is the function
// where migrations should be registered.
func Initialize() {
	log.Println("Initializing database migrations...")

	_ = goose.SetDialect("postgres")

	goose.AddNamedMigration("1_create_users_table.go", CreateUsersTableUp, CreateUsersTableDown)
	goose.AddNamedMigration("2_create_groups_info_table.go", CreateGroupsTableUp, CreateGroupsTableDown)
}

// Migrate parses environment for necessary parameters and starts
// database migration process.
func Migrate(db *sql.DB) {
	log.Println("Starting database migration procedure...")

	var action = "UP"

	actionFromEnv, actionFound := os.LookupEnv("DATABASE_ACTION")
	if actionFound {
		log.Println("Migration action override: " + actionFromEnv)
		action = actionFromEnv
	} else {
		log.Println("Executing default migration action (UP)")
	}

	var count int64

	countFromEnv, countFound := os.LookupEnv("DATABASE_COUNT")
	if countFound {
		log.Println("Migration count override: " + countFromEnv)

		countAsInt, err := strconv.ParseInt(countFromEnv, 10, 64)
		if err != nil {
			log.Fatalln("Failed to convert count gathered from DATABASE_COUNT to integer")
		}

		count = countAsInt
	} else {
		log.Println("Applying or rollback this count of migrations: " + countFromEnv + ". 0 - all.")
	}

	// Execute migrations.
	var err error

	currentDBVersion, gooseerr := goose.GetDBVersion(db)

	if gooseerr != nil {
		log.Fatalln("Failed to get database version: " + gooseerr.Error())
	}

	log.Println("Current database version obtained: " + strconv.Itoa(int(currentDBVersion)))

	switch {
	case action == "UP" && count == 0:
		log.Println("Applying all unapplied migrations...")

		err = goose.Up(db, ".")
	case action == "UP" && count != 0:
		newVersion := currentDBVersion + count
		log.Println("Migrating database to specific version: " + strconv.Itoa(int(newVersion)))
		err = goose.UpTo(db, ".", newVersion)
	case action == "DOWN" && count == 0:
		log.Println("Downgrading database to zero state, you'll need to re-apply migrations!")

		_ = goose.Down(db, ".")

		log.Fatalln("Database downgraded to zero state. You have to re-apply migrations")
	case action == "DOWN" && count != 0:
		newVersion := currentDBVersion - count
		log.Println("Downgrading database to specific version: " + strconv.Itoa(int(newVersion)))
		err = goose.DownTo(db, ".", newVersion)
	default:
		log.Fatalln("Unsupported set of migration parameters, cannot continue: " + action + "/" + countFromEnv)
	}

	if err != nil {
		log.Fatalln("Failed to execute migration sequence: " + err.Error())
	}

	log.Println("Database migrated successfully")
}
