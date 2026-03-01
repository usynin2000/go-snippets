package main

import (
	"fmt"
	"os"
)

func readFile(path string) error {
	_, err := os.ReadFile(path)
	if err != nil {
		// add context and wrap the original error
		return fmt.Errorf("error reading file %s: %w", path, err)
	}
	return nil
}

func main() {
	err := readFile("notfound.txt")
	if err != nil {
		fmt.Println("An error occurred:", err)
	}
}
// An error occurred: error reading file notfound.txt: open notfound.txt: no such file or directory
