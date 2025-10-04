// interface{} (в Go 1.18+ можно писать any) — это пустой интерфейс, который реализует любой тип.

// Часто используется для работы с "любым" значением.

package main

import "fmt"

func PrintAny(x interface{}) {
	fmt.Println(x)
}

func PrintAny2(x any) {
	fmt.Println(x)
}

func main() {
	PrintAny(42)
	PrintAny("hello")
	PrintAny([]int{1, 2, 3})
	PrintAny2(42)
	PrintAny2("hello")
	PrintAny2([]int{1, 2, 3})
}
