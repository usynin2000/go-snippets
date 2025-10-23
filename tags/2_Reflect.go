package main

import (
	"fmt"
	"reflect"
)

type Product struct {
	ID		int		`json:"id" db:"product_id"`
	Title	string	`json:"title" db:"product_name"`
}

func main() {
	t := reflect.TypeOf(Product{})

	lenFields := t.NumField()
	fmt.Println("Product length fields:", lenFields)

	for i := 0; i < lenFields; i++ {
		field := t.Field(i)
		fmt.Printf(
			"%s: json=%q db=%q\n",
			field.Name,
			field.Tag.Get("json"),
			field.Tag.Get("db"),
		)
	}
}