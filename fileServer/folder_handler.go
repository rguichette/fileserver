package fileserver

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FolderHandler struct {
	StorageDir string
}

func NewFileHandler(storageDir string) *FolderHandler {
	err := os.MkdirAll(storageDir, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Failed to create storage directory: %v", err))
	}
	return &FolderHandler{StorageDir: storageDir}
}

func (f *FolderHandler) Savefile(fileName string, content io.Reader) error {
	fullPath := filepath.Join(f.StorageDir, fileName)

	err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)

	if err != nil {
		return fmt.Errorf("failed to create directories:%w", err)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	if err != nil {
		return fmt.Errorf("failed to write file content: %w", err)
	}
	fmt.Printf("file saved: ;%s\n", fullPath)
	return nil

}
