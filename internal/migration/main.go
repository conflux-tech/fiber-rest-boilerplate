package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/conflux-tech/fiber-rest-boilerplate/configs"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

const dialect = "postgres"

func newDB(conf *configs.DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v", conf.Host, conf.Port, conf.User, conf.Database, conf.Password, conf.SSLMode)
	fmt.Println(connStr)
	return sql.Open(dialect, connStr)
}

var (
	flags = flag.NewFlagSet("migrate", flag.ExitOnError)
	dir   = flag.String("dir", "./database/migrations", "directory with migration files")
)

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])

	args := flags.Args()
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		flags.Usage()
		return
	}
	command := args[0]
	switch command {
	case "create":
		if err := goose.Run("create", nil, *dir, args[1:]...); err != nil {
			log.Fatalf("migrate run: %v", err)
		}
		return
	case "fix":
		if err := goose.Run("fix", nil, *dir); err != nil {
			log.Fatalf("migrate run: %v", err)
		}
		return
	}

	conf, _ := configs.Load()
	db, err := newDB(&conf.Database)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	if err := goose.SetDialect(dialect); err != nil {
		log.Fatal(err)
	}
	if err := goose.Run(command, db, *dir, args[1:]...); err != nil {
		log.Fatalf("migrate run: %v", err)
	}
}

func usage() {
	fmt.Println(usagePrefix)
	flags.PrintDefaults()
	fmt.Println(usageCommands)
}

var (
	usagePrefix = `Usage: migrate [OPTIONS] COMMAND
Examples:
    migrate status
Options:
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
    fix                  Apply sequential ordering to migrations
`
)
