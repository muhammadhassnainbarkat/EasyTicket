package main

import (
	"context"
	"flag"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log"
	"os"
)

var (
	flags = flag.NewFlagSet("migrate", flag.ExitOnError)
	dir   = flags.String("dir", "migrations", "directory with migration files")
)

func main() {
	flags.Usage = func() {
		fmt.Println(usagePrefix)
		flags.PrintDefaults()
		fmt.Println(usageCommands)
	}

	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("Error parsing flags: %v", err)
		return
	}
	args := flags.Args()
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		flags.Usage()
		return
	}

	command := args[0]

	dbString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Karachi",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))

	db, err := goose.OpenDBWithDriver("pgx", dbString)

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Error closing database: %v", err)
		}
	}()
	if err := goose.RunContext(context.Background(), command, db, *dir, args[1:]...); err != nil {
		log.Fatalf("migrate %v: %v", command, err)
	}

	fmt.Println("Migration Completed")

}

var (
	usagePrefix = `Usage: migrate COMMAND
Examples:
    migrate status
`

	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-by-one            Migrate the DB up by 1
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                  Apply sequential ordering to migrations`
)
