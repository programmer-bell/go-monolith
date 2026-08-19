package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/programmer-bell/go-monolith/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: make migrate <up | down>")
	}

	ctg := config.MustLoad()

	m, err := migrate.New(
		"file://migrations",
		ctg.DatabaseUrl,
	)
	if err != nil {
		log.Fatalf("migration.new:%v", err)
	}
	defer m.Close()

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown command: %s", os.Args[1])
	}

	fmt.Println("running migrations")
}
