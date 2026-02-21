package main

import "fmt"

// Equal reports whether a and b are the same length and contain the same bytes.

func Equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	a := []byte{'a', 'c'}
	b := []byte{'a', 'c'}
	c := []byte{'a', 'b', 'c'}

	isEqual_a_b := Equal(a, b)
	fmt.Println(isEqual_a_b)

	isEqual_a_c := Equal(a, c)
	fmt.Println(isEqual_a_c)



}
