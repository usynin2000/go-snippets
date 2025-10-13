package main

import (
	"fmt"
	"os"
)

func main() {
	envList := os.Environ()
	// выводим первые 10 элементов
	for i := 0; i < 10 && i < len(envList); i++ {
		fmt.Println(envList[i])
	}
}