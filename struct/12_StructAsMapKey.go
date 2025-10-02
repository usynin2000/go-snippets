package main


import "fmt"

type Point struct {
	X, Y int
}


func main() {
	m := make(map[Point]string)

	m[Point{1, 2}] = "A"
	m[Point{3, 4}] = "B"

	fmt.Println(m[Point{1, 2}]) // A

}