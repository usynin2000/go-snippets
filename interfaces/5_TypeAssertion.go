// Если у нас есть интерфейс, можно узнать, что внутри:
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// утверждение типа
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