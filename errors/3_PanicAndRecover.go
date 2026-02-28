// panic stops the program (like an exception).
// recover allows to catch a panic and continue execution.

package main

import "fmt"

func mayPanic() {
	panic("something went wrong")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Before panic")
	mayPanic()
	fmt.Println("After panic (not executed without recover)")
}

// Before panic
// Recovered: something went wrong


// When panic is called in mayPanic(),
// execution of the current function (main) is interrupted and the stack is unwound.
// At each level, deferred (defer) functions are executed.

// ⚠️ Important moment: recover() does not «return» the program back to the line after mayPanic().
// It only stops the program from falling.
// So execution continues after defer, not after mayPanic().

// If you want to continue the program after recover(),
// you need to wrap the «dangerous» code in a separate function.


// !!!! recover does not «return» you inside the function where the panic happened — it only stops the program from falling.
