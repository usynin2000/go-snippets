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
// Это встроенный роутер в Go, который умеет:

// "сопоставлять" URL-пути с обработчиками (handler’ами),

// и передавать запрос нужной функции.


// Зачем он нужен

// Если бы мы использовали только http.ListenAndServe(":8080", handler), то весь сервер обрабатывал бы одним обработчиком (как в твоём первом примере).

// Но когда нам нужно несколько страниц/эндпоинтов — нужен роутер, который решает:

// если путь / → вызвать mainPage,

// если путь /api → вызвать apiPage,

// если путь /static/... → вызвать обработчик для статики, и т. д.

// Вот этим и занимается ServeMux.


// Это нам это даёт

// Разделение логики — можно писать отдельные обработчики для разных страниц.

// Гибкость — можно подключить кастомные маршрутизаторы (например, Gorilla Mux) вместо стандартного.

// Читаемость — код явно показывает: /api → этот handler, / → другой handler.