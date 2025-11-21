package main


import "fmt"

func main() {
	// Anonymous function - function withont name
	func() {
		fmt.Println("Hello from anonymous function")
	}()

	// Anonymous function with params
	func(name string) {
		fmt.Println("Hello, %s!\n", name)
	}("Serega")

	// Assign anonymous function to variable
	greet := func(name string) {
		fmt.Println("Hello 2, %s\n", name)
	}
	
	greet("Zhopa")
	greet("Anya")

	// Anonymous func with returned values
	add := func(a, b int) int {
		return a + b
	}

	result := add(5, 3)
	fmt.Println("5 + 3 =", result)

	// Anonymous func as an argument to other func
	numbers := []int{1, 2, 3, 4, 5}
	processNumbers(
		numbers,
		func(n int) {
			fmt.Printf("Processing number: %d\n", n)
		},
	)


}

// function that recieve another func as param
func processNumbers(nums []int, processor func(int)) {
	for _, n := range nums {
		processor(n)
	}
}