// If you have interface, you can look what is inside
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// type assertion
	s, ok := i.(string)
	if ok {
		fmt.Println("String:", s)
	}

	val, ok := i.(int)
	if ok {
		fmt.Println("Int", val)
	} else {
		fmt.Println("It is not int")
	}
}
