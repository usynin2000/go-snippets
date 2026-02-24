// 🔹 Task 3. List of students

// Create a Student struct with fields Name and Grade.
// In main create a slice of students:

// Alice – 5, Bob – 4, Charlie – 3


// Iterate over the slice and print:

// Alice has grade 5
// Bob has grade 4
// Charlie has grade 3

// 🔹 Average grade

// Use the same Student struct and slice from task 3.
// Write a function Average(students []Student) float64 that calculates the average grade.
// Print its result.

// Expected output:

package main

import "fmt"

type Student struct {
	Name string
	Grade int
}

func Average(students []Student) float64 {
	var grades_together int
	for _, student := range students {
		grades_together += student.Grade
	}
	avg := float64(grades_together) / float64(len(students))
	return avg
}

func main() {
	students := []Student{{"Alice", 5}, {"Bob", 4}, {"Charlie", 3}}

	for _, student := range students {
		fmt.Println(student.Name, "has grade", student.Grade)
	}

	average := Average(students)
	fmt.Println("Average grade:", average)
}