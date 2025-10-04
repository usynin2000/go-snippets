// 🔹 Задание 3. Список студентов

// Создай структуру Student с полями Name и Grade (оценка).
// В main создай срез студентов:

// Alice – 5, Bob – 4, Charlie – 3


// Пройди по срезу и выведи:

// Alice has grade 5
// Bob has grade 4
// Charlie has grade 3

// 🔹 Средний балл

// Используй ту же структуру Student и срез из задания 3.
// Напиши функцию Average(students []Student) float64, которая считает среднюю оценку.
// Выведи её результат.

// Ожидаемый вывод:

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