package main

import (
	"encoding/json"
	"net/http"
)

type Subj struct {
	Product string `json:"name"`
	Price int `json:"price"`
}

func JSONHandler(w http.ResponseWriter, req *http.Request) {
	// take data
	subj := Subj{"Milk", 50}

	// encode to json
	resp, err := json.Marshal(subj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Set header Content-Type
	w.Header().Set("content-type", "application/json")
	// set status code 200
	w.WriteHeader(http.StatusOK)
	// write to the body of response
	w.Write(resp)
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", JSONHandler)

	err := http.ListenAndServe(":8080", m)
	if err != nil {
		panic(err)
	}
}
