// Создай структуру Rectangle с полями W и H (ширина и высота).
// Напиши метод Area(), который возвращает площадь.
// В main создай прямоугольник 10x5 и выведи его площадь.

// Ожидаемый вывод:
package main

import "fmt"

type Rectangle struct { W, H int}

func (r Rectangle) Area() int { return r.H * r.W}

func main() {
	rec := Rectangle{10, 5}

	res := rec.Area()

	fmt.Println(res)
}