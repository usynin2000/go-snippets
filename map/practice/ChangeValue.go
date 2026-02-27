// 5. Change value

// Create map[string]string with translations:
// "cat" → "cat", "dog" → "dog".
// Change translation "dog" to "dog".
package main

import "fmt"

func main() {
	m := map[string]string{
		"cat": "кот",
		"dog": "собака",
	}
	m["dog"] = "пёс"

	fmt.Println(m)
}
