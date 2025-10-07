package main

import (
	"net/http"
	"log"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://google.com", http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/", redirect)
	log.Fatal(http.ListenAndServe(":8080", nil))
}