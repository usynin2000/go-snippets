package main


import (
	"os"
	"fmt"
)

func main() {
	// first argument - name of the running file
	fmt.Printf("Command: %v/n", os.Args[0])

	// print other arguments
	for i, v := range os.Args[1:] {
		fmt.Println(i+1, v)
	}
}

// go run 1_GetArgs.go -thumb -w 1024 -file "./in/img0001.jpg" -dest "/home/user/images"