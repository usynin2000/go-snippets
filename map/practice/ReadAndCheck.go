// 2. Чтение и проверка

// Дан map[string]int{"Alice": 25, "Bob": 30}.
// Попробуй получить возраст "Charlie".
// Сначала просто выведи результат.
// А потом сделай через проверку val, ok := m["Charlie"].

package main

import "fmt"

func main() {
	m := map[string]int{"Alice": 25, "Bob": 30}

	fmt.Println(m["Charlie"])

	val, ok := m["Charlie"]
	if ok {
		fmt.Println("Charlise есть, вот возраст его", val)
	} else {
		fmt.Println("Мы не знаем кто такой Чарли и сколько ему лет.")
	}

}