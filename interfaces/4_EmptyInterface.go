package main

import "fmt"

// from GO version 1.18 + we can use any,
// interface{} == any

func PrintAny(x interface{}) {
	fmt.Println(x)
}

func PrintAny2(x any) {
	fmt.Println(x)
}

func main() {
	PrintAny(42)
	PrintAny("hello")
	PrintAny([]int{1, 2, 3})
	PrintAny2(42)
	PrintAny2("hello")
	PrintAny2([]int{1, 2, 3})
}
