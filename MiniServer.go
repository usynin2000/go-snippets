package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


type RequestData struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type ResponseData struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST поддерживается", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения запроса", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var data RequestData

	err = json.Unmarshal(body, &data)

	if err != nil {
		http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
	}

	resp := ResponseData{
		Message : fmt.Sprintf("Привет %s. Тебе %d лет.", data.Name, data.Age),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/hello", handler)
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}