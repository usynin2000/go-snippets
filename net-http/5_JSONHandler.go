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
	//  собираем данные
	subj := Subj{"Milk", 50}

	// кодируем в json
	resp, err := json.Marshal(subj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return 
	}

	// устанавливаем заголовок Content-Type
	// для передачи клиенту информации, кодированный в JSON
	w.Header().Set("content-type", "application/json")
	// устанавливаем код 200
	w.WriteHeader(http.StatusOK)
	// пишем тело ответа
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