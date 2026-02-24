package main

import "fmt"

type Counter struct { Value int}


// When we change structure -> user pointer receiver
func (c *Counter) Inc() {c.Value++}

// When we read -> value receiver
func (c Counter) Get() int { return c.Value}


func main() {

	c := Counter{0}

	pointerC := &c

	pointerC.Inc()
	pointerC.Inc()
	pointerC.Inc()

	fmt.Println(c.Get())

}
// Value receiver (func (r Rectange))
	// Using when method doesn't change structure
	// Wroks okay for small sctrucre (no need to copy much data)


// Pointer receiver (func (r *Rectangle))
	// Using  if method change strucre
	// Save space for big struct
