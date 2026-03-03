# 📝 Шпаргалка по структурам в Go

1. Объявление и инициализация
```go
type Person struct {
    Name string
    Age  int
}

p1 := Person{Name: "Alice", Age: 25} // явное указание полей
p2 := Person{"Bob", 30}              // по порядку (лучше избегать)
p3 := Person{}                       // все значения по умолчанию

```

2. Доступ и изменение
```go
p1.Age = 26
fmt.Println(p1.Name, p1.Age) // Alice 26
```

3. Анонимные структуры
```go
p := struct {
    Name string
    Age int
}{"Alice", 25}

fmt.Println(p) // {Alice 25}
```

4. Вложенные структуры
```go
type Address struct {
    City string
    Zip  string
}

type Person struct {
    Name string
    Addr Address
}

p := Person{
    Name: "Alice",
    Addr: Address{"Moscow", "101000"},
}

fmt.Println(p.Addr.City) // Moscow

```

5. Встраивание (аналог наследования)
```go
type Contact struct {
    Email string
    Phone string
}

type Employee struct {
    Name string
    Contact // встроенная структура
}

e := Employee{
    Name:    "Bob",
    Contact: Contact{"bob@mail.com", "+123"},
}

fmt.Println(e.Email) // bob@mail.com (можно обращаться напрямую)

```

6. Методы для структур
Методы объявляются глобально, а не внутри main.
```go
type Rectangle struct {
    W, H int
}

// метод-значение
func (r Rectangle) Area() int {
    return r.W * r.H
}

// метод-указатель
func (r *Rectangle) Scale(k int) {
    r.W *= k
    r.H *= k
}

r := Rectangle{5, 3}
fmt.Println(r.Area()) // 15
r.Scale(2)
fmt.Println(r.Area()) // 30

```

7. Сравнение структур
```go
type Point struct{ X, Y int }

p1 := Point{1, 2}
p2 := Point{1, 2}
p3 := Point{2, 3}

fmt.Println(p1 == p2) // true
fmt.Println(p1 == p3) // false

```
> ⚠️ Сравнивать можно только если все поля — сравнимые типы.


8. Структуры и срезы / map
```go
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 25},
    {"Bob", 30},
}

m := map[string]Person{
    "a": {"Charlie", 40},
}

fmt.Println(people[0].Name) // Alice
fmt.Println(m["a"].Age)     // 40

```

9. JSON с тегами
```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

p := Person{"Alice", 25}
data, _ := json.Marshal(p)
fmt.Println(string(data)) // {"name":"Alice","age":25}

```

10. Анонимные поля
```go
type User struct {
    string
    int
}

u := User{"Alice", 25}
fmt.Println(u.string, u.int) // Alice 25

```


11. Optional/Nullable поля через указатели
Для полей, которые могут быть "не заданы":
```go
type Update struct {
    Name *string  // nil = не обновлять, &"x" = обновить
    Age  *int     // nil = не обновлять, &0 = установить 0
}
```
Всегда проверяйте != nil перед разыменованием!
