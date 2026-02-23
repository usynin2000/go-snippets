package main

import "fmt"

func increment(n *int) {
	*n++
}

func main() {
	value := 10
	increment(&value)
	fmt.Println("After increment:", value)
}

// It helps to save memory and gives an ability to change var value from outer scope
