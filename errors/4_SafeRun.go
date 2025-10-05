package main

import "fmt"

func safeRun() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Восстановилось после:", r)
		}
	}()

	fmt.Println("Выполняем ...")
	panic("фатальная ошибка")
}

func main() {
	safeRun()
	fmt.Println("Программа работает дальше")
}