package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Example 1: Without Gorutines ===")
	// Ordinary func calling -> program awaits completion

	slowFunction("Func 1")
	slowFunction("Func 2")

	fmt.Println("Ordinary funcs stopped\n")

	fmt.Println("=== Example 2: With Gorutines ===")
	// Calling funcs with using of gorutines -> they will work in concurenncy 
	go slowFunction("Gorutine 1")
	go slowFunction("Gorutine 2")

	// Just to wait that gorutines completed
	time.Sleep(2 * time.Second)
	fmt.Println("Gorutines completed in parallel\n")

	fmt.Println("=== Example 3: Problem without wating ===")
	// NB! main func can be finished earlier that gorutines
	go slowFunction("Gorutine 3")
	fmt.Println("main is done, however we are not sure that gorutine is finished")
}


func slowFunction(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: шаг: %d\n", name, i)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Printf("%s: завершена\n", name)
}
