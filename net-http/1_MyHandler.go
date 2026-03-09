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


// Once MyHandler has a method with such a signature, it automatically becomes http.Handler.
// Why exactly an empty structure?

// Because we don't need to store any fields.

// We simply use MyHandler as a "carrier of the method".

// Alternative
// There is another, shorter way in Go:
// http.ListenAndServe(
// 	":8080",
// 	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//     w.Write([]byte("Привет!"))
// }))

// Here we don't create a structure, but directly use a function wrapper (http.HandlerFunc),
// which transforms a regular function into http.Handler.
