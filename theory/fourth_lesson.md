# Methods

В Go действительно нет классов, но можно определять методы для типов.

Что такое метод в Go?
Метод — это обычная функция, но с приёмником (receiver), который указывается перед именем функции. Этот приёмник определяет, к какому типу относится метод.
```go
package main

import "fmt"

// Определяем структуру Person
type Person struct {
    Name string
}

// Определяем метод SayHello для Person
func (p Person) SayHello() {
    fmt.Println("Привет, меня зовут", p.Name)
}

func main() {
    p := Person{"Алекс"}  // Создаём объект
    p.SayHello()          // Вызываем метод
}
```

Разбор кода:
Создаём структуру Person, у которой есть одно поле — Name (имя человека).
Создаём метод SayHello(), который:
Принимает p Person — это наш приёмник (receiver).
Выводит в консоль строку "Привет, меня зовут <имя>".
В main() создаём объект Person{"Алекс"}.
Вызываем метод SayHello(), и он выводит

Привет, меня зовут Алекс


> Разница в том, что в Go мы явно указываем приёмник метода (p Person), а в Python используется self.

Reminder:
В Go struct (структура) — это набор полей (переменных), объединённых в один тип. Это что-то похожее на класс в Python, но без методов и наследования.


Важные различия:
- struct фиксирован по типу – нельзя случайно записать строку в поле, которое ожидает число.
- Нет встроенных методов – но можно добавлять их вручную (как в примере с SayHello).
- Нет наследования – в отличие от классов Python.


Важно, что метод в GO можно писать и как просто функцию:
```go
package main

import "fmt"

// Определяем структуру Person
type Person struct {
    Name string
}

// Определяем метод SayHello для Person
func SayHello(p Person) {
    fmt.Println("Привет, меня зовут", p.Name)
}

func main() {
    p := Person{"Алекс"}  // Создаём объект
    SayHello(p)          // Вызываем метод
}
```


You can declare a method on non-struct types, too.

In this example we see a numeric type MyFloat with an Abs method.

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
```go
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

```


### Метод с указателем в качестве приемника
Когда мы передаём значение (value receiver), то работаем с копией объекта.
Когда мы передаём указатель (pointer receiver), то работаем с самим объектом и можем его изменить.


✨ Простой пример:
Допустим, у нас есть структура Person, и мы хотим изменить возраст с помощью метода.

🔴 Ошибочный вариант (без указателя, просто копия):
```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// Метод БЕЗ указателя (копия структуры)
func (p Person) Birthday() {
    p.Age++ // Увеличиваем возраст
}

func main() {
    p := Person{"Алекс", 30}
    p.Birthday() // Но изменения НЕ сохранятся!
    fmt.Println(p.Age) // 30 (НЕ изменилось!)
}

```

✅ Правильный вариант (используем указатель *Person):
```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// Метод С указателем (меняет оригинал)
func (p *Person) Birthday() {
    p.Age++ // Теперь изменяем оригинальный объект
}

func main() {
    p := Person{"Алекс", 30}
    p.Birthday() // Изменения сохранятся!
    fmt.Println(p.Age) // 31 (всё работает!)
}

```


💡 Что поменялось?

Мы написали func (p *Person) Birthday(), т.е. p теперь указатель на Person, а не копия.
Теперь метод Birthday() изменяет оригинальный объект, а не временную копию.



📌 Вывод:
> Если метод НЕ изменяет объект → передаём значение (Person).
> Если метод изменяет объект → передаём указатель (*Person).
> Так работает не только с struct, но и с любыми типами в Go. 🎯


### Почему так работает?
🔎 Как Go работает с переменными?
Когда ты передаёшь значение (p Person), Go копирует его.
Когда ты передаёшь указатель (p *Person), Go передаёт адрес в памяти (ссылку на оригинальный объект).

🔴 Передача по значению (без * — создаётся копия)
``` go
func (p Person) Birthday() {
    p.Age++  // Мы изменяем ТОЛЬКО копию
}
```
⏩ Что происходит?

p.Birthday() вызывает метод.
Go создаёт копию Person, новую переменную.
p.Age++ увеличивает возраст у копии, но оригинальный объект не изменяется.
Когда метод заканчивается, копия пропадает, а оригинальный объект остаётся без изменений.
✅ Передача по указателю (с * — работаем с оригиналом)
```go
func (p *Person) Birthday() {
    p.Age++  // Изменяем оригинальный объект
}
```
⏩ Что происходит?

p.Birthday() передаёт в метод указатель на Person (ссылку на объект в памяти).
p.Age++ изменяет тот же самый объект, который хранится в main().
Изменения сохраняются, потому что работаем не с копией, а с оригиналом.
🎯 Представь ситуацию в реальной жизни:
Без указателя (копия):

Ты дал другу копию карты сокровищ.
Он нарисовал новый путь, но на своей копии.
Твоя карта осталась неизменной.
С указателем (оригинал):

Ты дал другу свою единственную карту.
Он нарисовал новый путь прямо на ней.
Теперь карта изменилась и у тебя тоже.
💡 Вывод:

Без * → работаем с копией объекта → изменения не сохраняются.
С * → работаем с оригиналом → изменения сохраняются.



### В Go методы с указателем могут вызываться даже на значениях (структурах без *), и Go автоматически преобразует их в указатели.

🔴 Функции требуют явного указателя
```go
package main

import "fmt"

type Vertex struct {
    X, Y float64
}

// Обычная функция, принимающая УКАЗАТЕЛЬ
func ScaleFunc(v *Vertex, factor float64) {
    v.X *= factor
    v.Y *= factor
}

func main() {
    var v Vertex
    ScaleFunc(v, 5)  // ❌ Ошибка: ожидается указатель
    ScaleFunc(&v, 5) // ✅ Работает, передаём адрес (&v)
}
```

Почему ScaleFunc(v, 5) не работает?

Потому что ScaleFunc() требует указатель (*Vertex), но мы передаём значение (Vertex).
Go не делает автоматическое преобразование для функций.


✅ Методы с указателем работают и с объектами, и с указателями
```go
package main

import "fmt"

type Vertex struct {
    X, Y float64
}

// Метод Scale с УКАЗАТЕЛЕМ
func (v *Vertex) Scale(factor float64) {
    v.X *= factor
    v.Y *= factor
}

func main() {
    var v Vertex

    v.Scale(5)  // ✅ Работает, Go автоматически делает (&v).Scale(5)
    p := &v
    p.Scale(10) // ✅ Работает, т.к. p уже указатель
}
```
Почему v.Scale(5) работает, а ScaleFunc(v, 5) — нет?

> Для методов Go автоматически разыменовывает объект (v → &v).
Для функций такого механизма нет, и указатель нужно передавать вручную.


📌 Вывод:
> Если метод принимает *T, можно вызывать его на T, и Go сам сделает &T.
> Функции так не работают, указатель нужно передавать вручную.


### Когда использовать указатель (*T) в качестве получателя (receiver)?
Go предлагает два основных случая, когда методы должны использовать указатель (*T) вместо значения (T):

1. Метод должен изменять структуру
Если метод изменяет поля структуры, нужно использовать указатель *T, иначе изменения не сохранятся.

Пример:
```go
type Vertex struct {
    X, Y float64
}

// Метод изменяет объект → используем *Vertex
func (v *Vertex) Scale(factor float64) {
    v.X *= factor
    v.Y *= factor
}

func main() {
    v := Vertex{3, 4}
    v.Scale(2) // Эквивалентно (&v).Scale(2)
    fmt.Println(v) // {6, 8}
}
```

Если бы Scale принимал Vertex (без *), то работала бы копия структуры, а не оригинальный объект.


2. Избежать копирования при вызове метода
Если структура большая, передача копии может быть дорогой по памяти.

Пример:
```go
type LargeStruct struct {
    Data [1000000]byte // 1MB данных
}

// Используем *LargeStruct, чтобы не копировать 1MB при каждом вызове
func (ls *LargeStruct) Process() {
    fmt.Println("Processing large struct")
}
```
Если бы мы использовали Process(ls LargeStruct), то каждый вызов копировал бы 1MB, что замедляет программу.


Почему нельзя смешивать методы с T и *T?
Если часть методов использует T, а другая часть *T, то это может создать путаницу.
Go не всегда может преобразовать T в *T (и наоборот), особенно в интерфейсах.

Общий совет: если хотя бы один метод требует *T, лучше делать все методы *T.


## Interfaces
Если коротко, интерфейс в Go — это тип, который описывает набор методов. Любой тип, который реализует эти методы, автоматически удовлетворяет интерфейсу (в отличие от Python, где часто используется "duck typing" — если у объекта есть нужный метод, значит можно вызывать).

В Go интерфейсы — это способ абстракции. Они позволяют писать код, который может работать с разными типами данных, если эти типы реализуют нужный интерфейс.

Пример на Python (для контраста):
В Python ты бы написал что-то такое:
```python
class Dog:
    def speak(self):
        print("Woof!")

class Cat:
    def speak(self):
        print("Meow!")

def animal_sound(animal):
    animal.speak()

dog = Dog()
cat = Cat()

animal_sound(dog)  # Woof!
animal_sound(cat)  # Meow!
```
Здесь у нас нет интерфейсов, но функция animal_sound ожидает, что у объекта есть метод speak.


Как это выглядит в Go:
```go
package main

import "fmt"

// Интерфейс
type Speaker interface {
	Speak()
}

// Структура Dog
type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

// Структура Cat
type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

// Функция, принимающая интерфейс
func animalSound(s Speaker) {
	s.Speak()
}

func main() {
	dog := Dog{}
	cat := Cat{}

	animalSound(dog) // Woof!
	animalSound(cat) // Meow!
}

```

Что здесь важно:
Интерфейс описывает поведение, а не данные: type Speaker interface { Speak() }.

> Dog и Cat автоматически реализуют интерфейс Speaker, потому что у них есть метод Speak().

> Функция animalSound принимает любой тип, который удовлетворяет интерфейсу Speaker.


Главное отличие от Python:

В Python ты можешь передать в функцию любой объект, у которого есть нужный метод (динамическая типизация).
> В Go функция явно ожидает тип интерфейса, и компилятор проверяет, что переданный тип этот интерфейс реализует (строгая статическая типизация).

Когда это полезно:
Когда ты хочешь писать общий код для разных типов.
Для реализации полиморфизма (разные типы, но общее поведение).
Когда хочешь разделить зависимости (например, передать в функцию объект, который умеет логировать, независимо от того, что это за объект).


В Go, если тип имеет все методы, которые описаны в интерфейсе — он автоматически реализует этот интерфейс.
Тебе не нужно писать ничего вроде class Dog implements Speaker, как в Java или C#.

> Главная идея: "Если оно выглядит как утка, плавает как утка и крякает как утка — значит, это утка."


📜 Пример, чтобы это прочувствовать:
```go
type Speaker interface {
    Speak()
}
```

Структура (тип), который реализует метод:
```go
type Human struct{}

func (h Human) Speak() {
    fmt.Println("Hello")
}
```
И всё! 🚀 Human теперь реализует интерфейс Speaker, автоматически, просто потому что у него есть метод Speak().

## Почему это круто:
- Меньше зависимостей: Тип и интерфейс могут находиться даже в разных пакетах, и им не нужно знать друг о друге напрямую.
- Проще писать тесты: Можно описать интерфейс для нужного поведения и в тестах подсовывать mock-объекты.
- Гибкость: Легко добавлять новые реализации интерфейса без изменения самого интерфейса


### Что такое "interface value"?
Когда ты создаешь переменную типа интерфейса и присваиваешь туда конкретный тип, Go под капотом сохраняет две вещи:

Значение (сам объект, например, Dog{}).
Тип этого значения (например, Dog).
Именно поэтому интерфейсы в Go можно представить как кортеж (value, type).


## 💡 Что это значит на практике?
Когда ты вызываешь метод интерфейса, Go смотрит на конкретный тип, лежащий внутри интерфейса, и вызывает соответствующий метод.
```go
package main

import "fmt"

// Интерфейс
type Speaker interface {
	Speak()
}

// Структура
type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

func main() {
	var s Speaker   // Переменная типа интерфейса
	s = Dog{}       // Присваиваем конкретный тип (Dog) интерфейсу

	s.Speak()       // Вызывает Dog.Speak(), т.к. внутри лежит Dog
}

```

✅ Что под капотом у переменной s?

s = (Dog{}, тип Dog)


И когда вызывается s.Speak(), Go:

Видит, что внутри лежит Dog{}.
Ищет метод Speak у Dog.
Вызывает Dog.Speak().


🧠 Почему это важно?
Потому что интерфейс не знает заранее, какой там тип, но знает, что этот тип обязан иметь методы интерфейса.
И поведение интерфейса зависит от того, какой конкретный тип лежит внутри.

🔥 Пример с несколькими типами

```go
type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

func main() {
	var s Speaker

	s = Dog{}
	s.Speak() // Woof!

	s = Cat{}
	s.Speak() // Meow!
}

```

Что происходит?
Сначала s = (Dog{}, тип Dog) → вызывает Dog.Speak().
Потом s = (Cat{}, тип Cat) → вызывает Cat.Speak().


⚙️ Почему это "tuple of (value, type)"?
Чтобы Go знал, какой именно метод вызывать. Потому что интерфейсы могут принимать любой тип, который "совпадает" по методам, но вызывать нужно именно тот метод, который реализовал конкретный тип.


### Интерфейсы с nil внутри
📌 В чем суть?
В Go интерфейсная переменная (interface value) — это, напомню, (type, value).
И вот тут важный момент:

*Если в интерфейсе лежит тип + nil значение (например, ( SomeType, nil )), то сам интерфейс — НЕ nil!

🔑 Ключевые моменты:
Интерфейс может содержать nil-значение внутри, но сам интерфейс не равен nil, если внутри указан тип!
Если вызвать метод у такого интерфейса, то метод вызовется с nil receiver.
В Go часто пишут методы, которые умеют работать с nil receiver (в отличие от других языков, где будет NullPointerException).


🧠 Разберем на примере:
Пример структуры и интерфейса:
```go
package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}

func (d *Dog) Speak() {
	if d == nil {
		fmt.Println("Nothing to say (nil dog)")
		return
	}
	fmt.Println("Woof!")
}

func main() {
	var d *Dog = nil  // nil указатель на Dog

	var s Speaker = d // интерфейс содержит (тип *Dog, значение nil)

	fmt.Println(s == nil) // false — интерфейс сам по себе не nil!

	s.Speak() // вызовется метод, который обработает nil receiver
}

```

 Что тут происходит:
d — это *Dog, равный nil.

s — это интерфейс Speaker, внутри которого лежит nil, но тип указан:

s = (тип *Dog, значение nil)

s == nil → false! Потому что тип есть, даже если значение nil.

При вызове s.Speak() вызывается метод Speak() для типа *Dog, и внутрь передается nil, но метод обрабатывает это спокойно:
Nothing to say (nil dog)


🎯 Как правильно проверить на "пустой интерфейс"?

Нужно проверить оба элемента (тип и значение). Обычно это делается через:
```go
if s == nil || reflect.ValueOf(s).IsNil() { ... }
```

⚡ Контраст с Python:
В Python:
```python
x = None
if x is None:
    print("None")

```

Но в Go:
```go
var d *Dog = nil
var s Speaker = d // НЕ nil интерфейс

if s == nil {
    fmt.Println("Nil interface") // НЕ выполнится
}

```

🚀 Резюме:
Ситуация	Результат
Интерфейс пустой  (не содержит типа и значения)	 -> nil
Интерфейс содержит тип + nil значение	-> НЕ nil, но значение внутри nil
Вызов метода у интерфейса с nil значением внутри	 -> Метод вызывается, получает nil как receiver


Как избежать сюрпризов?
> 	Проверять не только интерфейс, но и значение внутри


### The empty interface
The interface type that specifies zero methods is known as the empty interface:

interface{}

An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.

```go
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

```


### Type assertions
A type assertion provides access to an interface value's underlying concrete value.

t := i.(T)

This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

If i does not hold a T, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

t, ok := i.(T)
If i holds a T, then t will be the underlying value and ok will be true.

If not, ok will be false and t will be the zero value of type T, and no panic occurs.

Note the similarity between this syntax and that of reading from a map.

```go
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
```

OUTPUT:
hello
hello true
0 false
panic: interface conversion: interface {} is string, not float64

goroutine 1 [running]:
main.main()
	/tmp/sandbox4188577700/prog.go:17 +0x13f



## Type switches
Type switches
A type switch is a construct that permits several type assertions in series.

A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.

This switch statement tests whether the interface value i holds a value of type T or S. In each of the T and S cases, the variable v will be of type T or S respectively and hold the value held by i. In the default case (where there is no match), the variable v is of the same interface type and value as i.

```go
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```


## Stringers
One of the most ubiquitous interfaces is Stringer defined by the fmt package.

type Stringer interface {
    String() string
}
A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
```


### Errors
Go programs express error state with error values.

The error type is a built-in interface similar to fmt.Stringer:

type error interface {
    Error() string
}
As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
> A nil error denotes success; a non-nil error denotes failure.

21/26