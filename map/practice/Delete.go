// 3. Удаление

// Дан map[string]int{"Alice": 25, "Bob": 30, "Charlie": 40}.
// Удалить "Bob".
// Выведи карту.
package main

import "fmt"

func main() {
	m := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 40}

	delete(m, "Bob")

	fmt.Println(m)
}