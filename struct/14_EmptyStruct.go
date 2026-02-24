package main

import "fmt"


func main() {
	set := make(map[string]struct{}) // This creates map (hash table), where keys are string, and values are empty struct struct{}.

	set["apple"] = struct{}{}
	set["banana"] = struct{}{}

	// We insert two elements in the "set": "apple" and "banana".

	_, ok := set["apple"]
	fmt.Println("Is there apple(key)?", ok) // true

}


// Why exactly map[string]struct{}

// In Go there is no type Set (like in python)

// But we can make "Set" by using map.

// Key (string) is element of our "Set"

// Value (struct{}) is just simple dummy

// Can we useap[string]bool?

// Yes we can. But struct{} is more cheap.

// struct{} as empty structure does not use any memory (- bite)

// bool use 1 byte for every element

// That is why the right Go-way for using set in map[T]struct{}


