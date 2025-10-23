package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name	string	`json:"name"` // имя поля в JSON
	Age		int		`json:"age,omitempty` // пропускается, если 0
	Email	string 	`json:"-"` //  исключить поле из JSON
}


func main() {
	u := User{Name: "Alice", Age: 0, Email: "usynin.s00@mail.ru"}
	data, _ := json.Marshal(u) // сделает json  но в байтах [123 34 110 97 109 101 34 58]
	fmt.Println(string(data))
}
