package main

import "fmt"

type Counter struct {
	Value int
}

func (c *Counter) Inc() {
	c.Value++
}

func main() {
	c := Counter{10}
	c.Inc()
	fmt.Println(c.Value) // 11
}
