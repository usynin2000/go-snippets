package main

import (
	"net/http"
)

func main() {
	// simple server which can use all file in dir statis
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
	if err != nil {
		panic(err)
	}
}
