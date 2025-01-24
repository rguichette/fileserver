package websockets

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all origins (will adjust as needed)
	},
}

func StartWebSocketServer(address string) error {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			fmt.Printf("Failed to upgrade connection: %v\n", err)
			return
		}

		fmt.Println("Websocket connection established")
		handleWebSocketConnection(conn)

	})

	//start the server
	return http.ListenAndServe(address, nil)

}

func handleWebSocketConnection(conn *websocket.Conn) {
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			break
		}

		fmt.Printf("Received message: %s\n", message)

		//echo the messaage mack to client

		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Printf("Error writing message: %v\n", err)
			break
		}
	}

}
