## 📘 Интерфейсы в Go
1. Определение

Интерфейс в Go — это набор методов, которые должен реализовать тип.
Если тип реализует все методы интерфейса, то он автоматически подходит под этот интерфейс.
```go
type Shape interface {
    Area() float64
}
```
2. Реализация интерфейса
```go
type Rectangle struct {
    W, H float64
}

// метод для Rectangle
func (r Rectangle) Area() float64 {
    return r.W * r.H
}

func main() {
    var s Shape
    s = Rectangle{3, 4}
    fmt.Println(s.Area()) // 12
}
```

> ➡️ Rectangle реализует Shape, потому что у него есть метод Area.


## 3. Зачем нужны интерфейсы

- Абстракция: код работает не с конкретными типами, а с поведением.

- Полиморфизм: разные типы можно использовать одинаково.

- Снижение связности: функции зависят не от типа, а от набора методов.


### 4. Интерфейсы в стандартной библиотеке

error — любой тип с методом Error() string.

fmt.Stringer — любой тип с методом String() string.

io.Reader, io.Writer — стандартные интерфейсы для работы с потоками.
```go
type MyError struct{}

func (e MyError) Error() string {
    return "oops"
}

func main() {
    var err error = MyError{}
    fmt.Println(err.Error()) // oops
}
```

5. Пустой интерфейс (interface{} / any)

interface{} — это интерфейс без методов.
➡️ Любой тип автоматически подходит под него.
```go
func PrintAny(x interface{}) {
    fmt.Println(x)
}

func main() {
    PrintAny(42)
    PrintAny("Hello")
    PrintAny(true)
}

```

6. Извлечение значения

Когда мы принимаем interface{}, нужно понять, какой тип внутри.
Type assertion
```go
var i interface{} = "hi"
s, ok := i.(string) // извлекаем строку
fmt.Println(s, ok)  // hi true

```

Type switch
```go
switch v := i.(type) {
case int:
    fmt.Println("int:", v)
case string:
    fmt.Println("string:", v)
default:
    fmt.Println("unknown")
}

```

7. Важные моменты

В Go нет ключевого слова implements.
👉 Реализация интерфейса происходит автоматически.

Интерфейсы используют:

- как контракты методов (например, io.Reader, fmt.Stringer),

- как универсальные контейнеры (any, interface{}).