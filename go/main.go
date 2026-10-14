package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	addr := net.JoinHostPort("::", port)

	log.Printf("Listening on %s", addr)

	if err := http.ListenAndServe(addr, routes()); err != nil {
		log.Fatal(err)
	}
}
