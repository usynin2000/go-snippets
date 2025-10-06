package main

import (
	"fmt"
	"net/http"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)
	body += "Header =================\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	res.Write([]byte(body))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", mainPage)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

// запусти http://localhost:8080/?id=12345&name=John%20Doe&filter=town&filter=country.