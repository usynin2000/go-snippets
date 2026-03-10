package main

import (
	"io"
	"fmt"
	"net/http"
)

func WriteHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "1")
	fmt.Fprint(w, "2")
	w.Write([]byte("3"))
}

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(WriteHandler))
	if err != nil {
		panic(err)
	}
}

//  123

// It's important to understand that http.ResponseWriter in Go is an interface,
// which itself extends another interface:

// Тут важно понять, что http.ResponseWriter в Go — это интерфейс,
// который сам по себе расширяет другой интерфейс:
// type ResponseWriter interface {
//     Header() Header
//     Write([]byte) (int, error)   // реализует io.Writer
//     WriteHeader(statusCode int)
// }

