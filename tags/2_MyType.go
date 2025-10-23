package main

import (
	"fmt"
	"reflect"
	"time"
)

type MyType struct {
	User      string    `json:"user,omitempty" example:"Bob"` // Здесь у поля User есть два тега: json и example.
	CreatedAt time.Time `json:"created_at"`
}

const (
	targetField = "User" // имя поля, о котором нужно получить информацию
	targetTag   = "json" // тег, значение которго нужно получить
)

func main() {

	obj := MyType{}

	// получаем GO-описание типа
	objType := reflect.TypeOf(obj) // Получение информации о типе

	fmt.Println("objType", objType) // main.MyType  
	// объект типа reflect.Type, который содержит метаданные о типе MyType.

	// ищем поле по имени
	field, ok := objType.FieldByName(targetField)
	// FieldByName() находит поле структуры по его имени и возвращает reflect.StructField.
	if !ok {
		panic(fmt.Errorf("field (%s): not found", targetField))
	}

	fmt.Println("field:", field) // 

	// ищем тег по имени
	tagValue, ok := field.Tag.Lookup(targetTag) // Tag.Lookup() извлекает значение конкретного тега из поля.
	if !ok {
		panic(fmt.Errorf("tag (%s) for field (%s): not found", targetTag, targetField))
	}

	fmt.Printf("Значение тега %s поля %s: %s\n", targetTag, targetField, tagValue)
	fmt.Printf("Теги поля %s: %s\n", targetField, field.Tag)
}
