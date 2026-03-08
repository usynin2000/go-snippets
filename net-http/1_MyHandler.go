package main

import "net/http"

// We are creating the empy sttruct MyHandler
// It will be implement interface http.Handler
type MyHandler struct{}

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Привет!")
	res.Write(data)
}
// Struct MyHander has method ServerHTTP
// This is method is required to implement interfact http.Handler
// Args:
// - res http.ResponseWriter - object, which we use to write response for client (sending data)
// - req *http.Request — object with data about request (with method, headers and path)


func main() {
	var h MyHandler

	err := http.ListenAndServe(`:8080`, h)
	if err != nil {
		panic(err)
	}
}


// That how interface http.Handler lookds like
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
