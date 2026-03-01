package main

import (
	"errors"
	"fmt"
	"os"
)

func openFile(path string) error {
	_, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	return nil
}

func main() {
	err := openFile("notfound.txt")

	// Check if it is exactly "file not found"
	if errors.Is(err, os.ErrNotExist){
		fmt.Println("File not found")
	}

	// Check through errors.As (for example *os.PathError)
	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		fmt.Println("Error associated with file:", pathErr.Path)
	}
}

// File not found
// Error associated with file: notfound.txt
