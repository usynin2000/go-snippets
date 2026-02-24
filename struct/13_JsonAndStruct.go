package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

// `...` is tag
// It can be read by differnt libs
// It is pretty useful because in GO it is common to use CamelCase, and snake_case in JSON

func main() {
	p := Person{"Alice", 25}

	data, _ := json.Marshal(p)
	fmt.Println(string(data)) // {"name":"Alice","age":25}
}
