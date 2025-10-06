package main

import "net/http"

// Создаётся пустая структура MyHandler. 
// Она будет реализовывать интерфейс http.Handler.
type MyHandler struct{}

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Привет!")
	res.Write(data)
}
// У структуры MyHandler определён метод ServeHTTP.
// Этот метод обязателен для реализации интерфейса http.Handler.
// Аргументы:
// - res http.ResponseWriter — объект, через который мы пишем ответ клиенту (отправляем данные).
// - req *http.Request — объект с информацией о запросе (метод, заголовки, путь и т.д.).
// Внутри:
// - Создаём []byte("Привет!").
// - res.Write(data) записывает эти данные в ответ. В итоге браузер/клиент получит строку "Привет!".


func main() {
	var h MyHandler

	err := http.ListenAndServe(`:8080`, h)
	if err != nil {
		panic(err)
	}
}


// В этом коде:

// type MyHandler struct{}
// мы завели её только ради того, чтобы:

// На неё "повесить" метод ServeHTTP.

// У Go нет "свободных" методов — метод всегда должен принадлежать какому-то типу.

// Поэтому мы создаём тип (MyHandler) и делаем для него метод ServeHTTP.

// Чтобы этот тип соответствовал интерфейсу http.Handler.

// Интерфейс http.Handler выглядит так:
// type Handler interface {
//     ServeHTTP(ResponseWriter, *Request)
// }


// Как только у MyHandler появляется метод с такой сигнатурой — он автоматически считается http.Handler.


// Но почему именно пустая структура?

// Потому что нам не нужно хранить никаких полей.

// Мы просто используем MyHandler как "носитель метода".

// Можно было бы и так:
// type MyHandler int
// или
// type MyHandler string
// — лишь бы метод можно было привязать. 
// Но пустая структура логичнее: она ничего не весит и явно показывает, что данные нам не нужны.


// Альтернатива
// В Go есть и другой, более короткий способ:
// http.ListenAndServe(
// 	":8080",
// 	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//     w.Write([]byte("Привет!"))
// }))

// Здесь мы не создаём структуру, а сразу используем функцию-обёртку (http.HandlerFunc), 
// которая превращает обычную функцию в http.Handler.