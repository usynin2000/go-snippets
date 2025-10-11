package main


import (
	"os"
	"fmt"
)

func main() {
	// первый аргумнт - имя запущенного файла
	fmt.Printf("Command: %v/n", os.Args[0])

	// выведем остальные параметры
	for i, v := range os.Args[1:] {
		fmt.Println(i+1, v)
	}
}

// go run 1_GetArgs.go -thumb -w 1024 -file "./in/img0001.jpg" -dest "/home/user/images"