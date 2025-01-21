package fileserver_test

import (
	"testing"
	//
	fileserver "github.com/rguichette/tcplib/fileServer"
)

func TestTCPServerStart(t *testing.T) {
	handler := fileserver.SimpleHandler{}

	//create a new TCPServer instance
	server := fileserver.NewTCPServer(":8090", &handler)

	//Attempting to start the server
	err := server.Start()
	if err != nil {
		t.Fatalf("Failed to start the server: %v", err)
	}

	if server.Listener == nil {
		t.Fatalf("Expected a non-nil listener after starting the server")
	}

	err = server.Start()
	if err != nil {
		t.Fatalf("Server failed to handle repeated start: %v", err)
	}

	_ = server.Listener.Close()

}

func TestTCPServerStop(t *testing.T) {
	handler := fileserver.SimpleHandler{}
	server := fileserver.NewTCPServer(":8090", &handler)

	err := server.Start()
	if err != nil {
		t.Fatalf("Failed to start the server: %v", err)
	}

	err = server.Stop()
	if err != nil {
		t.Fatalf("Failed to stop the server: %v", err)
	}

	//make sure the server listener is nil after stopping
	if server.Listener != nil {
		t.Fatalf("Expected the listener to be nil after stopping the server")
	}

	err = server.Stop()
	if err != nil {
		t.Fatalf("Server failed to handle repeated stop: %v", err)
	}
}
