package main

import "fmt"

type MyError struct {
	Code int
	Msg string
}

func (e MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Msg)
}

func doSomething(fail bool) error {
	if fail {
		return MyError{404, "Not found"}
	}
	return nil
}

func main() {
	err := doSomething(true)
	if err != nil {
		fmt.Println(err)
	}
}
