package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var bearer = "Beater <Token>"

func main() {
	// Создаем новый запрос
	req, err := http.NewRequest("GET", "https://yandex.ru", nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Добавляем авторизацию
	req.Header.Add("Authorization", bearer)

	// создаем клиента
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Bad status code on response: ", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	// продолжаем работу
	fmt.Println(string(body))



}