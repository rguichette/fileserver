package fileserver

import (
	"fmt"
	"net"
	"sync"
)

type Server interface {
	Start()
	Stop()
}
type ConnectionHandler interface {
	Handle(conn net.Conn)
}

type SimpleHandler struct{}

func (h *SimpleHandler) Handle(conn net.Conn) {
	defer conn.Close()

	fmt.Fprintln(conn, "Welcome to the TCP server!")
	fmt.Println("Handled a connection")
}

type TCPServer struct {
	address  string
	Listener net.Listener
	running  bool

	handler ConnectionHandler
	mu      sync.Mutex
}

func (s *TCPServer) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.running {
		return nil
	}

	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	s.Listener = listener
	s.running = true

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				if s.running {
					fmt.Printf("Error accepting connection: %v\n", err)
				}
				return
			}

			fmt.Println("New connection accepted")
			// conn.Close()
			// continue
			go s.handler.Handle(conn)

		}
	}()
	return nil
}

func (s *TCPServer) Stop() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.running {
		return nil
	}

	err := s.Listener.Close()
	if err != nil {
		return err
	}
	s.Listener = nil
	s.running = false
	return nil
}

func NewTCPServer(address string, handler ConnectionHandler) *TCPServer {
	return &TCPServer{
		address: address,
		handler: handler,
	}
}
