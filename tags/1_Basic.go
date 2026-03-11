package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name	string	`json:"name"` // field name in JSON
	Age		int		`json:"age,omitempty"` // skipped if 0 or not passed
	Email	string 	`json:"-"` // exclude field from JSON
}


func main() {
	u := User{Name: "Serega", Age: 0, Email: "usynin.s00@mail.ru"}
	data, _ := json.Marshal(u)
	fmt.Println(string(data))
}
