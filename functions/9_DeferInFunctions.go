package main

import "fmt"

// defer - postpone function execution to ending of fuction

func example1() {
	fmt.Println("Beginnig of fuction.")
	defer fmt.Println("postponed call N 1")
	defer fmt.Println("postponed call N 2")
	fmt.Println("The end of fuction")
}

func example2(name string) {
	defer fmt.Println("Function is completed.")
	fmt.Printf("Handling: %s\n", name)
	if name == "" {
		return // defer still going to be completed
	}
	fmt.Println("Continue handling")
}

func example3() {
	x := 1
	defer fmt.Println("x in defer", x) // x will be 1 (because in the moment of defer it was 1)
	x = 2
	fmt.Println("x after changing", x)
}

func main() {
	fmt.Println("=== Example 1 ===")
	example1()

	fmt.Println("\n=== Example 2 ===")
	example2("Test")
	example2("")

	fmt.Println("\n=== Example 3 ===")
	example3()

}
