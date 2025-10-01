package main

import "fmt"

type Person struct {
	Name string
	Age int
}

type Employee struct {
	Person // встраивание
	Position string
}

func main() {
	 e := Employee {
		Person: Person{Name: "Alice", Age: 25},
		Position: "Developer",
	 }

	 fmt.Println(e.Name) // доступ к полю Person напрямую
	 fmt.Println(e.Position)  // свое поле 
}

// Это один из способов композиции: когда одна структура "встраивается" внутрь другой, и её поля и методы становятся доступны напрямую,
//  как будто они принадлежат внешней структуре.


// Если бы ты написал так:
// type Employee struct {
//     P        Person
//     Position string
// }

// То для доступа к имени пришлось бы писать:
// e.P.Name


// Вместо имени поля (P) ты пишешь просто тип:
// type Employee struct {
//     Person   // без имени
//     Position string
// }

// Теперь Go "поднимает" поля и методы Person на уровень Employee.
// Поэтому:
// fmt.Println(e.Name)   // работает
// fmt.Println(e.Age)    // тоже работает