package main

import (
	"log"

	fileserver "github.com/rguichette/tcplib/fileServer"
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

	storageDir := "./storage"
	folderHandler := fileserver.NewFileHandler(storageDir)
	handler := &fileserver.SimpleHandler{FolderHandler: folderHandler}

	server := fileserver.NewTCPServer(":8090", handler)

	log.Println("Starting server on :8090")
	err := server.Start()

	if err != nil {
		log.Fatal("Failed to start the server: %v", err)
	}

	select {}

	// fileName := "example/testfile.txt"
	// fileConent := strings.NewReader("this is a test file")

	// err := handler.Savefile(fileName, fileConent)
	// if err != nil {
	// 	panic(err)
	// }

	// println("file saved successfule, check the strdir")

}
