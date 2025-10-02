package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	p := Person{"Alice", 25}

	data, _ := json.Marshal(p)
	fmt.Println(string(data)) // {"name":"Alice","age":25}
}

// То, что в обратных кавычках (`...`), называется тегом структуры.

// Тег — это строка, привязанная к полю, и её могут читать пакеты рефлексии (reflect), например encoding/json.

// Для JSON-тегов:

// json:"name" — говорит, что поле Name будет сериализовано как "name", а не "Name".
// Это удобно, так как в Go принято CamelCase, а в JSON — snake_case или lowerCamelCase.

// json:"age" — аналогично для возраста.

// Если тег не указать:
// type Person struct {
//     Name string
//     Age  int
// }

// То результат будет:
// {"Name":"Alice","Age":25}

// Можно ещё управлять сериализацией:

// json:"name,omitempty" — если поле пустое (нулевое значение), оно не попадёт в JSON.

// json:"-" — поле полностью игнорируется при сериализации/десериализации.
