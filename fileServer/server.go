package fileserver

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type Server interface {
	Start()
	Stop()
}
type ConnectionHandler interface {
	Handle(conn net.Conn)
}

type SimpleHandler struct {
	FolderHandler *FolderHandler
}

func (h *SimpleHandler) handlerListCommand(conn net.Conn) {
	files, err := h.FolderHandler.ListFiles()
	if err != nil {
		fmt.Fprintf(conn, "Error: %v\n", err)
		return
	}

	for _, file := range files {
		fmt.Fprintln(conn, file)
	}
	fmt.Fprintln(conn, "END OF LIST")
}

func (h *SimpleHandler) Handle(conn net.Conn) {
	fmt.Fprintln(conn, "Welcome to the TCP server!.. inside handle")
	fmt.Println("Handled a connection")

	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	var currentFile string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "LIST") {
			h.handlerListCommand(conn)
		} else if strings.HasPrefix(line, "UPLOAD") {
			currentFile = strings.TrimPrefix(line, "UPLOAD")
			fmt.Fprintf(conn, "Ready to receive %s\n ", currentFile)
		} else if currentFile != "" {
			content := strings.NewReader(line)
			err := h.FolderHandler.Savefile(currentFile, content)
			if err != nil {
				fmt.Fprint(conn, "Error saving successfully", err)
				currentFile = ""
			}

		} else {
			fmt.Fprintln(conn, "Error: invalid command")
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading connection: %v\n", err)
	}

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

			fmt.Println("New connection accepted -test")
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
