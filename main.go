package main

import (
	"log"

	fileserver "github.com/rguichette/tcplib/fileServer"
)

func main() {
	handler := &fileserver.SimpleHandler{}
	server := fileserver.NewTCPServer(":8090", handler)

	log.Println("Starting server on :8090")
	err := server.Start()

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	select {}
}
