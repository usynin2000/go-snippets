
# NamedReturns
```go
package main

import "fmt"

// Named return values - можно указать имена возвращаемых значений
func divide(a, b float64) (quotient float64, remainder float64) {
	quotient = a / b
	remainder = a - (quotient * b)
	return // naked return - возвращает именованные значения
}

// Можно использовать только одно имя, если возвращается одно значение
func calculate(x, y int) (result int) {
	result = x * y
	result += 10
	return // автоматически вернет result
}

// Именованные возвращаемые значения полезны для документации
func splitName(fullName string) (firstName, lastName string) {
	// В Go можно использовать naked return для именованных значений
	parts := []rune(fullName)
	mid := len(parts) / 2
	firstName = string(parts[:mid])
	lastName = string(parts[mid:])
	return
}

func main() {
	q, r := divide(10.0, 3.0)
	fmt.Printf("10 / 3 = %.2f, remainder: %.2f\n", q, r)

	res := calculate(5, 3)
	fmt.Println("calculate(5, 3):", res)

	first, last := splitName("ИванПетров")
	fmt.Printf("First: %s, Last: %s\n", first, last)
}

```


# 4_AnonymousFunctions
```go
package main

import "fmt"

func main() {
	// Анонимная функция - функция без имени
	func() {
		fmt.Println("Привет из анонимной функции!")
	}() // () вызывает функцию сразу

	// Анонимная функция с параметрами
	func(name string) {
		fmt.Printf("Привет, %s!\n", name)
	}("Иван")

	// Присваивание анонимной функции переменной
	greet := func(name string) {
		fmt.Printf("Здравствуй, %s!\n", name)
	}
	greet("Мария")
	greet("Петр")

	// Анонимная функция с возвращаемым значением
	add := func(a, b int) int {
		return a + b
	}
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)

	// Анонимная функция как аргумент другой функции
	numbers := []int{1, 2, 3, 4, 5}
	processNumbers(numbers, func(n int) {
		fmt.Printf("Обрабатываю число: %d\n", n)
	})
}

// Функция, принимающая другую функцию как параметр
func processNumbers(nums []int, processor func(int)) {
	for _, n := range nums {
		processor(n)
	}
}

```
# 5_Closures
```go
package main

import "fmt"

// Замыкание (closure) - функция, которая захватывает переменные из внешней области видимости
func makeCounter() func() int {
	count := 0 // переменная захватывается замыканием
	return func() int {
		count++ // изменяем захваченную переменную
		return count
	}
}

// Замыкание с параметром
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor // factor захватывается из внешней области
	}
}

// Замыкание для генерации последовательностей
func makeSequence(start, step int) func() int {
	current := start
	return func() int {
		value := current
		current += step
		return value
	}
}

func main() {
	// Создаем счетчик
	counter1 := makeCounter()
	fmt.Println("Counter 1:", counter1()) // 1
	fmt.Println("Counter 1:", counter1()) // 2
	fmt.Println("Counter 1:", counter1()) // 3

	// Другой счетчик - независимый
	counter2 := makeCounter()
	fmt.Println("Counter 2:", counter2()) // 1 (свой собственный count)

	// Множитель
	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	fmt.Println("double(5):", double(5))   // 10
	fmt.Println("triple(5):", triple(5))   // 15

	// Последовательность
	seq := makeSequence(10, 5)
	fmt.Println("Sequence:", seq()) // 10
	fmt.Println("Sequence:", seq()) // 15
	fmt.Println("Sequence:", seq()) // 20
}


```

# 6_Recursioin
```go
package main

import "fmt"

// Рекурсия - функция вызывает сама себя
func factorial(n int) int {
	if n <= 1 {
		return 1 // базовый случай
	}
	return n * factorial(n-1) // рекурсивный вызов
}

// Рекурсивный расчет чисел Фибоначчи
func fibonacci(n int) int {
	if n <= 1 {
		return n // базовые случаи: fib(0)=0, fib(1)=1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Рекурсивный поиск в слайсе
func binarySearch(arr []int, target, left, right int) int {
	if left > right {
		return -1 // не найдено
	}
	mid := (left + right) / 2
	if arr[mid] == target {
		return mid // найдено
	}
	if arr[mid] > target {
		return binarySearch(arr, target, left, mid-1) // ищем слева
	}
	return binarySearch(arr, target, mid+1, right) // ищем справа
}

// Рекурсивный обход дерева (простой пример)
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func traverseTree(node *Node) {
	if node == nil {
		return // базовый случай
	}
	fmt.Print(node.Value, " ")
	traverseTree(node.Left)  // рекурсивно обходим левое поддерево
	traverseTree(node.Right) // рекурсивно обходим правое поддерево
}

func main() {
	fmt.Println("Factorial(5):", factorial(5)) // 120

	fmt.Print("Fibonacci(0-9): ")
	for i := 0; i < 10; i++ {
		fmt.Print(fibonacci(i), " ")
	}
	fmt.Println()

	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	index := binarySearch(arr, 7, 0, len(arr)-1)
	fmt.Printf("Binary search for 7: index %d\n", index)

	// Простое дерево
	root := &Node{
		Value: 1,
		Left:  &Node{Value: 2, Left: &Node{Value: 4}, Right: &Node{Value: 5}},
		Right: &Node{Value: 3},
	}
	fmt.Print("Tree traversal: ")
	traverseTree(root)
	fmt.Println()
}


```


# 7_HigherOrderFunctions
```go
package main

import "fmt"

// Функции высшего порядка - функции, которые принимают или возвращают другие функции

// Map - применяет функцию к каждому элементу слайса
func mapInts(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter - фильтрует элементы слайса по условию
func filterInts(slice []int, fn func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce - сводит слайс к одному значению
func reduceInts(slice []int, initial int, fn func(int, int) int) int {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// Функция, возвращающая функцию
func makeValidator(min, max int) func(int) bool {
	return func(n int) bool {
		return n >= min && n <= max
	}
}

// Функция, принимающая несколько функций
func processWithPipeline(value int, functions ...func(int) int) int {
	result := value
	for _, fn := range functions {
		result = fn(result)
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Map: удваиваем каждое число
	doubled := mapInts(numbers, func(n int) int {
		return n * 2
	})
	fmt.Println("Doubled:", doubled)

	// Filter: оставляем только четные числа
	evens := filterInts(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Evens:", evens)

	// Reduce: суммируем все числа
	sum := reduceInts(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Println("Sum:", sum)

	// Валидатор
	isValidAge := makeValidator(18, 65)
	fmt.Println("Age 25 valid:", isValidAge(25))
	fmt.Println("Age 70 valid:", isValidAge(70))

	// Пайплайн функций
	addOne := func(n int) int { return n + 1 }
	multiplyByTwo := func(n int) int { return n * 2 }
	square := func(n int) int { return n * n }

	result := processWithPipeline(3, addOne, multiplyByTwo, square)
	fmt.Println("Pipeline(3):", result) // ((3+1)*2)^2 = 64
}


```


# 8_FunctionTypes.go
```go
package main

import "fmt"

// Определение типа функции
type MathOperation func(int, int) int
type Predicate func(int) bool
type Transformer func(int) int

// Функции можно использовать как типы
func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

// Функция, принимающая функцию как параметр
func calculate(a, b int, op MathOperation) int {
	return op(a, b)
}

// Массив функций
func applyOperations(value int, operations []Transformer) int {
	result := value
	for _, op := range operations {
		result = op(result)
	}
	return result
}

// Функция возвращает функцию
func getOperation(opType string) MathOperation {
	switch opType {
	case "add":
		return add
	case "multiply":
		return multiply
	default:
		return func(a, b int) int { return 0 }
	}
}

// Map с использованием типа функции
func mapSlice(slice []int, fn Transformer) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func main() {
	// Использование функций как значений
	var op MathOperation = add
	fmt.Println("add(5, 3):", op(5, 3))

	op = multiply
	fmt.Println("multiply(5, 3):", op(5, 3))

	// Передача функции в другую функцию
	result := calculate(10, 5, add)
	fmt.Println("calculate(10, 5, add):", result)

	result = calculate(10, 5, multiply)
	fmt.Println("calculate(10, 5, multiply):", result)

	// Получение функции из функции
	addOp := getOperation("add")
	fmt.Println("getOperation('add')(7, 3):", addOp(7, 3))

	// Массив операций
	operations := []Transformer{
		func(n int) int { return n + 1 },
		func(n int) int { return n * 2 },
		func(n int) int { return n - 5 },
	}
	final := applyOperations(10, operations)
	fmt.Println("applyOperations(10, ops):", final) // ((10+1)*2)-5 = 17

	// Использование mapSlice
	numbers := []int{1, 2, 3, 4, 5}
	squared := mapSlice(numbers, func(n int) int {
		return n * n
	})
	fmt.Println("Squared:", squared)
}


```

# 9_DeferinFunction.go
```go
package main

import "fmt"

// defer - отложенное выполнение функции до возврата из текущей функции

func example1() {
	fmt.Println("Начало функции")
	defer fmt.Println("Отложенный вызов 1")
	defer fmt.Println("Отложенный вызов 2")
	fmt.Println("Конец функции")
	// defer выполняются в обратном порядке (LIFO - Last In First Out)
}

func example2(name string) {
	defer fmt.Println("Функция завершена")
	fmt.Printf("Обрабатываю: %s\n", name)
	if name == "" {
		return // defer все равно выполнится
	}
	fmt.Println("Продолжаю обработку")
}

// defer с аргументами - аргументы вычисляются сразу, но выполнение откладывается
func example3() {
	x := 1
	defer fmt.Println("x в defer:", x) // x будет равен 1 (значение на момент defer)
	x = 2
	fmt.Println("x после изменения:", x)
}

// defer с замыканием - захватывает текущее значение
func example4() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println("i в замыкании:", i) // все выведут 3! (значение на момент возврата)
		}()
	}
}

// Правильный способ - передать значение как аргумент
func example5() {
	for i := 0; i < 3; i++ {
		defer func(n int) {
			fmt.Println("i через аргумент:", n) // выведет 2, 1, 0
		}(i) // значение i копируется в n
	}
}

// defer для очистки ресурсов
func processFile(filename string) error {
	fmt.Printf("Открываю файл: %s\n", filename)
	defer fmt.Printf("Закрываю файл: %s\n", filename) // всегда закроется
	
	fmt.Println("Обрабатываю файл...")
	if filename == "" {
		return fmt.Errorf("пустое имя файла") // defer выполнится перед return
	}
	fmt.Println("Файл обработан успешно")
	return nil
}

func main() {
	fmt.Println("=== Example 1 ===")
	example1()

	fmt.Println("\n=== Example 2 ===")
	example2("тест")
	example2("")

	fmt.Println("\n=== Example 3 ===")
	example3()

	fmt.Println("\n=== Example 4 (проблема) ===")
	example4()

	fmt.Println("\n=== Example 5 (решение) ===")
	example5()

	fmt.Println("\n=== Example 6 (очистка ресурсов) ===")
	processFile("data.txt")
	processFile("")
}


```
