package main

import (
	"os"

	"github.com/spplatform/edu-server/internal/server"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	server.Serve(host, port)
}
