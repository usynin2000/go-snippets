package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	City	string	`json:"city"`
	State	string	`json"state"`
}

type User struct {
	Name string `json:"name"`
	Address Address `json:"address"`
}

func printTags(t reflect.Type, prefix string) {
	fmt.Println("t", t)
	for i:= 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fieldName := prefix + f.Name

		if f.Type.Kind() == reflect.Struct {
			printTags(f.Type, fieldName+".") // рекурсия 
		} else {
			fmt.Printf("%s -> json:%q\n", fieldName, f.Tag.Get("json"))
		}
	}
}

func main() {
	t := reflect.TypeOf(User{})
	printTags(t, "")
}