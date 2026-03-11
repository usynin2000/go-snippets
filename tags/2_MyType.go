package main

import (
	"fmt"
	"reflect"
	"time"
)

type MyType struct {
	User      string    `json:"user,omitempty" example:"Bob"` // Here User field has two tags: json and example.
	CreatedAt time.Time `json:"created_at"`
}

const (
	targetField = "User" // name of field to get information about
	targetTag   = "json" // tag to get value of
)

func main() {

	obj := MyType{}

	// get GO-description of type
	objType := reflect.TypeOf(obj) // get information about type

	fmt.Println("objType", objType) // main.MyType

	// find field by name
	field, ok := objType.FieldByName(targetField)
	// FieldByName() finds field by name and returns reflect.StructField.
	if !ok {
		panic(fmt.Errorf("field (%s): not found", targetField))
	}

	fmt.Println("field:", field) //

	// find tag by name
	tagValue, ok := field.Tag.Lookup(targetTag) // Tag.Lookup() извлекает значение конкретного тега из поля.
	if !ok {
		panic(fmt.Errorf("tag (%s) for field (%s): not found", targetTag, targetField))
	}

	fmt.Printf("Value of tag %s of field %s: %s\n", targetTag, targetField, tagValue)
	fmt.Printf("Tags of field %s: %s\n", targetField, field.Tag)
}
