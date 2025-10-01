package main

import "fmt"

type Counter struct {
	Value int
}

func increase(c *Counter) {
	c.Value++
}

func main() {
	c := Counter{Value: 10}
	increase(&c)
	fmt.Println(c.Value) // 11
}