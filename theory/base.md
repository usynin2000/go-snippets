Programs start running in package main.

his program is using the packages with import paths "fmt" and "math/rand".

> By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.


```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

This code groups the imports into a parenthesized, "factored" import statement.

You can also write multiple import statements, like:

import "fmt"
import "math"
But it is good style to use the factored import statement.



## Exported names

In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

pizza and pi do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi) // math.pi не будет работать, 
}
```

То есть всегда капс после импортной функции.



### Типы в функциях:
Notice that the type comes after the variable name.
```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```


When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

In this example, we shortened
```go
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```



### Multiple results
A function can return any number of results.

The swap function returns two strings.
```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```


### Named return values
Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to **document the meaning of the return values.**

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```


### Variables
The var statement declares a list of variables; as in function argument lists, the type is last.

A var statement can be at package or function level. We see both in this example.


> The var statement declares a list of variables; as in function argument lists, the type is last.

A var statement can be at package or function level. We see both in this example.

```go
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}

```


A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```



### Short variable declarations
Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

> **Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.**
```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

## Basic types

Go's basic types are

- bool

- string

- int  int8  int16  int32  int64
- uint uint8 uint16 uint32 uint64 uintptr

- byte // alias for uint8

- rune // alias for int32
     // represents a Unicode code point

- float32 float64

- complex64 complex128

The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

```go
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

```shell
Type: bool Value: false
Type: uint64 Value: 18446744073709551615
Type: complex128 Value: (2+3i)
```


> Variables declared without an explicit initial value are given their zero value.

The zero value is:

- 0 for numeric types,
- false for the boolean type, and
- "" (the empty string) for strings.


### Type conversions
The expression T(v) converts the value v to the type T.

Some numeric conversions:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
Or, put more simply:

i := 42
f := float64(i)
u := uint(f)
Unlike in C, in Go assignment between items of different type requires an **explicit conversion.**

```go
var x, y int = 3, 4
var f float64 = math.Sqrt(float64(x*x + y*y))
```
> Без преобразования вышла бы ошибка


### Constants
Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.
```go
const Pi = 3.14
```