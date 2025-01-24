package main

import (
	"log"

	fileserver "github.com/rguichette/tcplib/fileServer"
	"github.com/rguichette/tcplib/fileServer/websockets"
)

func main() {
	// handler := &fileserver.SimpleHandler{}
	// server := fileserver.NewTCPServer(":8090", handler)

	// log.Println("Starting server on :8090")
	// err := server.Start()

	// if err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }
	// select {}

	// storageDir := "./storage"

	// //test when the directory does not exist
	// err := os.MkdirAll(storageDir, os.ModePerm)
	// if err != nil {
	// 	log.Fatalf("failed to create directory: %v ", err)
	// } else {
	// 	fmt.Println("directory created successfully (or already exists)")
	// }

	// //Test when the directory already exists
	// err = os.MkdirAll(storageDir, os.ModePerm)
	// if err != nil {
	// 	log.Fatalf("Failed to create directory again: %v", err)
	// } else {
	// 	fmt.Println("Directory already exists, no error occurred ")
	// }

	// storageDir := "./storage"
	// folderHandler := fileserver.NewFileHandler(storageDir)
	// handler := &fileserver.SimpleHandler{FolderHandler: folderHandler}

	// server := fileserver.NewTCPServer(":8090", handler)

	// log.Println("Starting server on :8090")
	// err := server.Start()

	// if err != nil {
	// 	log.Fatal("Failed to start the server: %v", err)
	// }

	// select {}

	// fileName := "example/testfile.txt"
	// fileConent := strings.NewReader("this is a test file")

	// err := handler.Savefile(fileName, fileConent)
	// if err != nil {
	// 	panic(err)
	// }

	// println("file saved successfule, check the strdir")

	/*ATTEMPTING to upgrade current tcp server to websocket*/

	handler := &fileserver.SimpleHandler{}
	tcpServer := fileserver.NewTCPServer(":8090", handler)

	go func() {
		log.Println("Starting TCP server on port 8090")
		if err := tcpServer.Start(); err != nil {
			log.Fatalf("Failed to start TCP server: %v", err)
		}
	}()

	// starting the websocket
	log.Println("starting websocket on port 8080")
	if err := websockets.StartWebSocketServer(":8080"); err != nil {
		log.Fatalf("Failed to start websocket server:%v", err)
	}

}
