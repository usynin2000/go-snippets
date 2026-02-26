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

// 💡 You can store slices inside structs and dynamically add elements.
