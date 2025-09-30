package main

import "fmt"

func main() {
	// Числа

	var i int = 42 // целое число (размер зависит от платформы)
	var i8 int8 = -128 // 8-битное целое (-128..127)
	var i16 int16 = 32000 // 16-битное целое
	var i32 int32 = 2_000_000 // 32-битное целое
	var i64 int64 = 9_000_000_000 // 64-битное целое

	var u uint = 42 // беззнаковое целое (>=0)
	var u8 uint8 = 255  // 8-битное беззнаковое (0..255)
	var u16 uint16 = 65000 // 16-битное беззнаковое
	var u32 uint32 = 4_000_000 // 32-битное беззнаковое
	var u64 uint64 = 18_000_000_000 // 64-битное беззнаковое

	var f32 float32 = 3.14 // 32-битное число с плавающей точкой
	var f64 float64 = 2.718281828   // 64-битное число (обычно используется чаще)
	
	var c64 complex64 = 1 + 2i  // комплексное число (float32)
	var c128 complex128 = 2 + 3i // комплексное число (float64)


	// Логический тип
	var b bool = true

	// Строки
	var s string = "Hello, Go!"

	// Символы
	var r rune = 'Я' // rune = int32 (UTF-8 символ)
	var by byte = 'A'  // byte = uint8 (один байт)

	// Указатели
	var p *int = &i

	// Массивы
	var arr[3]int = [3]int{1, 2, 3}

	// Срезы
	slice := []int{10, 20, 30}

	// Map
	m := map[string]int{"Alice": 25, "Bob": 30}

	// Struct
	type Person struct {
		Name string
		Age int
	}

	alice := Person{"Alice", 25}

	// Interface
	var any interface {} = "любое значение"

	// Функции 
	add := func(a, b int) int {
		return a + b
	}

	// Каналы
	ch := make(chan int)

	// ВЫВОД
	fmt.Println("Числа:", i, i8, i16, i32, i64, u, u8, u16, u32, u64, f32, f64)
	fmt.Println("Комлексные:", c64, c128)
	fmt.Println("Булев:", b)
	fmt.Println("Строка:", s)
	fmt.Println("Rune:", r, string(r))
	fmt.Println("Byte:", by, string(by))
	fmt.Println("Указатель:", p, *p)
	fmt.Println("Массив:", arr)
	fmt.Println("Срез:", slice)
	fmt.Println("Map:", m)
	fmt.Println("Struct:", alice)
	fmt.Println("Interface:", any)
	fmt.Println("Функция:", add(2, 3))
	fmt.Println("Канал:", ch)

}