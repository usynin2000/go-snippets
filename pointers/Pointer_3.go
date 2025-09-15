package main

import "fmt"

func increment(n *int) {
	*n++
}

func main() {
	value := 10
	increment(&value)
	fmt.Println("После increment:", value)
}

// 👉 Экономия памяти и возможность менять переменные "снаружи".