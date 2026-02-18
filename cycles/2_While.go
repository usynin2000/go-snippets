package main

import "fmt"

func main() {
	// While like cycle
	counter := 0
	for counter < 5 {
		fmt.Printf("while cycle: %d\n", counter)
		counter++
	}

	// Infinite loop with break condition
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Printf("Infinite cycle %d\n", i)
		i++
	}

	// Infinite loop withcontinue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("Odd numbers: %d\n", i)
	}

}
