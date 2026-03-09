package main

import "net/http"

func mainPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Привет!"))
}

func apiPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Это страница /api"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", apiPage)
	mux.HandleFunc("/", mainPage)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

// Что такое mux?
// ServeMux = multiplexer (мультиплексор).
// It's a built-in router in Go, which can:

// "сопоставлять" URL-пути с обработчиками (handler’ами),

// и передавать запрос нужной функции.


// Why it's needed?

// If we would use only http.ListenAndServe(":8080", handler), the whole server would be handled by one handler (as in your first example).

// When we need multiple pages/endpoints — we need a router, which solves:

// if the path is / → call mainPage,

// if the path is /api → call apiPage,

// if the path is /static/... → call the handler for static files, and so on.

// This is what ServeMux does.
