// Ошибки как возвращаемое значение
package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}

func main() {
	res, err := divide(10, 2)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результатэ:", res)
	}

	res, err = divide(10, 0)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", res)

	}
}