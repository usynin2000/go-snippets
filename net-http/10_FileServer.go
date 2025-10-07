package main

import (
	"net/http"
)

func main() {
	// простейшни сервер, которому доступны все файлы в поддиректории static
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
	if err != nil {
		panic(err)
	}
}