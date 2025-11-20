package main

import "fmt"


func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultVal V) V {
	if val, exists := m[key]; exists {
		return val
	}
	return defaultVal
}

func main() {
	users := map[string]string{
		"john": "John Doe",
		"jane": "Jane Smith",
	}

	fmt.Println(GetOrDefault(users, "john", "Unknown"))
	fmt.Println(GetOrDefault(users, "bob", "Unknown"))

	ages := map[string]int {
		"alice": 25,
		"bob": 30,
	}

	fmt.Println(GetOrDefault(ages, "alice", 0))
	fmt.Println(GetOrDefault(ages, "charlie", 0))

}