// 2. Срез и append

// Создай пустой срез []string, добавь туда имена "Alice", "Bob", "Charlie" и выведи длину и вместимость.

package main

import "fmt"

func main() {
	var s1 []string
	s1 = append(s1, "Alice", "Bob", "Charlie")

	fmt.Println(s1, len(s1), cap(s1))
}