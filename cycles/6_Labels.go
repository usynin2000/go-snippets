package main

import "fmt"

func main() {
	// Using labels to exit from nested loop
	fmt.Println("Search with using labels:")

outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1{
				fmt.Printf("Found position[%d][%d], exiting from all cycles \n", i, j)
				break outerLoop
			}
			fmt.Printf("Checking [%d][%d] \n", i, j)
		}
	}

	fmt.Println("\n Contitnue with label:")

continueLoop:
	for i := 0; i< 5; i++ {
		for j := 0; j < 3; j++ {
			if i == 2 && j == 1 {
				fmt.Printf("Skipping iteration [%d][%d]\n", i, j)
				continue continueLoop
			}
			fmt.Printf("Checking [%d][%d]\n", i, j)
		}
	}


}

