// Многие вещи в Go работают через интерфейсы:

// io.Reader, io.Writer → чтение/запись

// fmt.Stringer → метод String() string

// error → метод Error() string

// Пример с error:
package main

import "fmt"

type MyError struct{}

func (MyError) Error() string {
	return "something went wrong"
}

func main() {
	var err error = MyError{}
	fmt.Println(err.Error())
}