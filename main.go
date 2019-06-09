package main

import (
	"log"
	"os"

	"github.com/spplatform/edu-server/internal/migration"
	"github.com/spplatform/edu-server/internal/server"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	if len(os.Args) < 2 {
		return
	}

	switch os.Args[1] {
	case "migrate":
		migration.Migrate("migrations", "init")
		err := migration.Migrate("migrations", "up")
		if err != nil {
			log.Fatal(err)
		}
	case "serve":
		server.Serve(host, port)
	default:
		log.Fatal("false operation")
	}
}
