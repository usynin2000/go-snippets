package main

import "fmt"

// Stack with generic usage
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	// Stack of numbers
	stackInt := Stack[int]{}
	stackInt.Push(1)
	stackInt.Push(2)
	stackInt.Push(3)

	for !stackInt.IsEmpty() {
		val, _ := stackInt.Pop()
		fmt.Println(val)
	}

	// Stack of strings
	stackStr := Stack[string]{}
	stackStr.Push("Hello")
	stackStr.Push("World")
	val, _ := stackStr.Pop()
	fmt.Println(val)


}
