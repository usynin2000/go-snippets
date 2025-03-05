### Pointers

Go has pointers. A pointer **holds the memory address of a value.**

The type *T is a pointer to a T value. Its zero value is nil.

var p *int
The & operator generates a pointer to its operand.

i := 42
p = &i


The * operator denotes the pointer's underlying value.

fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p

This is known as "dereferencing" or "indirecting".

Unlike C, Go has no pointer arithmetic.
```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```

OUTPUT:
```shell
42
21
73
```

## Structs
A struct is a collection of fields.
```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

OUTPUT {1 2}

Struct fields are accessed using a dot.
```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
```

4/27