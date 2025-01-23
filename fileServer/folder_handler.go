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

func (f *FolderHandler) Readfile(fileName string) (string, error) {
	fullPath := filepath.Join(f.StorageDir, fileName)

	content, err := os.ReadFile(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(content), nil
}

func (f *FolderHandler) ListFiles() ([]string, error) {
	var files []string

	err := filepath.Walk(f.StorageDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}
		if !info.IsDir() {
			relPath, _ := filepath.Rel(f.StorageDir, path)
			files = append(files, relPath)
		}

		return nil
	})
	return files, err
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
