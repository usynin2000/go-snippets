package main

import (
	"net/http"
	"fmt"
)

func middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Serega-Ebet", "da")
		fmt.Println("In middleware")

		next.ServeHTTP(w, r)
	})
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ti gorish kak agonya"))
}

func main() {
	http.Handle("/", middleware(http.HandlerFunc(rootHandler)))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}