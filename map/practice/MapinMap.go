// 6. Вложенный map

// Создай map[string]map[string]int, где ключ — имя человека, а значение — карта "предмет → оценка".
// Например:

// {
//   "Alice": {"math": 5, "english": 4},
//   "Bob": {"math": 3, "english": 5}
// }


// Выведи оценку Боба по английскому.
package main

import "fmt"

func main() {
	m := map[string]map[string]int{
		"Alice": {"math": 5, "english": 4},
		"Bob": {"math": 3, "english": 5},
	}

	fmt.Println("Оценка боба по английскому", m["Bob"]["english"])
}