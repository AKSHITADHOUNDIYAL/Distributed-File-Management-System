package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Base directory where files are stored
const storageDir = "data/"

// Initialize storage directory
func InitStorage() error {
	if err := os.MkdirAll(storageDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create storage directory: %v", err)
	}
	return nil
}

// SaveFile stores a file in the worker's storage
func SaveFile(filename string, data io.Reader) error {
	filePath := filepath.Join(storageDir, filename)

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Copy data into file
	_, err = io.Copy(file, data)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	fmt.Println("File saved:", filePath)
	return nil
}

// GetFile retrieves a file from storage
func GetFile(filename string) (*os.File, error) {
	filePath := filepath.Join(storageDir, filename)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}

	return file, nil
}

// DeleteFile removes a file from storage
func DeleteFile(filename string) error {
	filePath := filepath.Join(storageDir, filename)

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	fmt.Println("File deleted:", filePath)
	return nil
}
