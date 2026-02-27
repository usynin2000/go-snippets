package main

import "fmt"


func main() {
	ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 40}

	for name, age :=range ages {
		fmt.Println(name, "-->", age)
	}
}
