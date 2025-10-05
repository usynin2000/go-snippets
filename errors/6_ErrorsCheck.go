package main

import (
	"errors"
	"fmt"
	"os"
)

func openFile(path string) error {
	_, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Не удалость открыть файл: %w", err)
	}
	return nil
}

func main() {
	err := openFile("notfound.txt")

	// Проверяем - именно ли "файл не найден"
	if errors.Is(err, os.ErrNotExist){
		fmt.Println("Файл отсутствует")
	}

	// Проверяем через errors.As (например *os.PathError)
	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		fmt.Println("Ошибка связана с файлом:", pathErr.Path)
	}
}