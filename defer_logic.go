package main

import (
	"fmt"
)



func main () {
	fmt.Println("Beginning of func")
	defer fmt.Println("This will be in the end of func")
	fmt.Println("Main code. Will be in the middle")
}

// Почему defer лучше, чем resp.Body.Close() в конце?
// Если json.NewDecoder().Decode() вызовет ошибку, до resp.Body.Close() код не дойдёт — соединение останется открытым!