package main

import "os"

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	Serve(host, port)
}
