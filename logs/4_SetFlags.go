package main

import (
	"log"
	"os"
)

func main() {
	// создаём файл info.log и обрабатываем ошибку, если что-то пошло не так
	file, err := os.OpenFile(
		"info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644,
	)
	if err != nil {
		log.Fatal(err)
	}

	// откладываем закрытие файла

	defer file.Close()

	log.SetOutput(file)

	 // добавляем флаги для вывода даты, времени, имени файла
	 log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	 log.Print("loggin in file GO")






}