package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// encoding/json — пакет для работы с JSON (кодирование и декодирование).
// fmt — стандартный пакет для вывода текста (например, в консоль).
// net/http — пакет для работы с HTTP (запросы, ответы, сервер).


type Response struct {
	Message string `json:"message"`
}

// Здесь мы создаём структуру Response, которая описывает JSON-объект.
// Message string — поле типа string, которое будет содержать сообщение.
// json:"message" — тэг, который указывает, как это поле будет называться в JSON.

// Пример JSON-ответа, который создаст эта структура:

func handler(w http.ResponseWriter, r *http.Request) {
	//w http.ResponseWriter — объект для отправки ответа клиенту.
	//r *http.Request — объект с данными запроса.


	// Устанавливаем заголовок, чтобы браузер понял, что это JSON
	w.Header().Set("Content-Type", "application/json")

	// Создаём объект структуры Response, который будет содержать текст ответа.
	response := Response{Message: "Привет, это Go-сервер!"}

	// Используем json.NewEncoder(w).Encode(response), чтобы преобразовать объект response в JSON и отправить его клиенту.	
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handler) // Устанавливаем обработчик для корневого маршрута
	// Говорим серверу: если поступает HTTP-запрос на адрес /, вызывай функцию handler.

	port := 8080
	fmt.Printf("Сервер запущен на порту %d...\n", port)
	// Printf (как f строки в python)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	// http.ListenAndServe(":8080", nil) — запускает сервер на порту 8080.
	// nil означает, что сервер использует обработчики, которые мы настроили ранее (http.HandleFunc).
}

