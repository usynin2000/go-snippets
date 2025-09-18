package main

import "fmt"

type Team struct {
	Name string
	Members []string
}

func main() {
	t := Team{
		Name: "Avengers",
		Members: []string{"Iron Man", "Thor"},
	}

	t.Members = append(t.Members, "Hulk")

	fmt.Println(t.Members)
}

// 💡 Можно хранить срезы внутри структур и динамически добавлять элементы.