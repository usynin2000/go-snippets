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
