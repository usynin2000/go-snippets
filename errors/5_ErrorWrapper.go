package main

import (
	"fmt"
	"os"
)

func readFile(path string) error {
	_, err := os.ReadFile(path)
	if err != nil {
		//  добавляем контекст и оборачиваем оргинальнюу ошибку 
		return fmt.Errorf("ошибка при чтении файла %s: %w", path, err)
	}
	return nil
}

func main() {
	err := readFile("notfound.txt")
	if err != nil {
		fmt.Println("Произошла ошибка:", err)
	}
}
// Произошла ошибка: ошибка при чтении файла notfound.txt: open notfound.txt: no such file or directory
