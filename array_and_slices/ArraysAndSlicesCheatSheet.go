package main

import "fmt"

func main() {
	// 1. Создание срезов
	s1 := []int{1, 2, 3} // литерал
	s2 := make([]int, 3, 5)
	fmt.Println("1) s1: ", s1, "| s2: ", s2, "len:", len(s2), "cap:", cap(s2))

	// 2.Slicing
	arr := []int{10, 20, 30, 40, 50}
	part1 := arr[1:3] // индксы 1 и 2
	part2 := arr[:2] // от начала до инрдекса 2
	part3 := arr[2:]  //  от 2 до конца
	fmt.Println("2) Slicing: ", part1, part2, part3)

	// 3. Добавление элемента
	s1 = append(s1, 4, 5)
	fmt.Println("3) append:", s1)

	// 4. Объединение срезов
	a := []int{1, 2}
	b := []int{3, 4, 5}
	a = append(a, b...)
	fmt.Println("4) Объединение:", a)

	// 5. Копирование среза
	src := []int{1, 2, 3}
	cp_src := make([]int, len(src))
	copy(cp_src, src)
	cp_src[0] = 100
	fmt.Println("5) copy: src =", src, "| cp_src =", cp_src)

	// 6. Удаление элемента
	nums := []int{1, 2, 3, 4, 5}
	index := 2
	nums = append(nums[:index], nums[index+1:]...)
	fmt.Println("6) Удаление элемента: ", nums)

	// 7. Iteration
	for i, v := range s1 {
		fmt.Println("7) range:", i, v)
	}

	// 8. nil и пустые срезы
	var sNil []int // равно nil, len == 0 
	sEmpty := []int{} // не равно nil, хотя len == 0
	fmt.Println("8) nil slice: ", sNil, "len:", len(sNil), "is_nil?", sNil == nil)
	fmt.Println("8) empty slice:", sEmpty, "len:", len(sNil), "is_nil?", sEmpty == nil)

	// 9. Filters
	names := []string{"Alice", "Bob", "Charlie", "Dave"}
	var short []string
	for _, name := range names {
		if len(name) <= 4 {
			short = append(short, name)
		}
	}
	fmt.Println("9) Фильтрация:", short)

	// 10. Подводный камень: разделение памяти
	orig := []int{1, 2, 3, 4, 5}
	alias := orig[:3] // указывает на тот же массив
	alias[0] = 99
	fmt.Println("10) aliasing:", "orig =", orig, "| alias =", alias)

	// Чтобы избежать: копируем
	independent := append([]int(nil), orig[:3]...)
	independent[0] = 123
	fmt.Println("   independent copy:", "orig =", orig, "| independent =", independent)

}