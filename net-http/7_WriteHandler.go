package main

import (
	"io"
	"fmt"
	"net/http"
)

func WriteHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "1")
	fmt.Fprint(w, "2")
	w.Write([]byte("3"))
}

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(WriteHandler))
	if err != nil {
		panic(err)
	}
}

// выведет 123

// Всё дело в том, что в Go все три вызова внутри твоего WriteHandle пишут данные в http.ResponseWriter.

// Тут важно понять, что http.ResponseWriter в Go — это интерфейс, 
// который сам по себе расширяет другой интерфейс:
// type ResponseWriter interface {
//     Header() Header
//     Write([]byte) (int, error)   // реализует io.Writer
//     WriteHeader(statusCode int)
// }


// Ключевая часть:
// 👉 У ResponseWriter есть метод Write([]byte), а значит он реализует интерфейс io.Writer.

// Теперь по порядку:

// io пакет

// Функции из io (например, io.WriteString) принимают любой io.Writer.

// Так как http.ResponseWriter реализует io.Writer, его можно туда передать.

// fmt пакет

// fmt.Fprint и его друзья (Fprintf, Fprintln) тоже пишут в любой io.Writer.

// Опять же, ResponseWriter подходит, потому что у него есть метод Write([]byte).

// Сам http.ResponseWriter

// Ты можешь напрямую вызвать w.Write([]byte(...)), потому что это базовый метод интерфейса.


// Пока твой объект (w) реализует интерфейс io.Writer, его можно скормить в любой пакет, который с ним умеет работать — будь то io, fmt, log, json.Encoder, gzip.Writer, и так далее.

// Это и есть сила композиции в Go:

// не нужны наследования,

// достаточно реализовать метод Write([]byte) → и ты уже можешь использовать тонну стандартных функций.