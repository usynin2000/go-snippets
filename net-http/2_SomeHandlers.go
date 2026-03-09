package main

import "net/http"

func mainPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello!"))
}

func apiPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("This page is /api."))
}

func main() {
	http.HandleFunc(`/api`, apiPage)
	http.HandleFunc(`/`, mainPage)

	err := http.ListenAndServe(`:8080`, nil)
	if err != nil{
		panic(err)
	}
}
