package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var bearer = "Beater <Token>"

func main() {
	// create new query
	req, err := http.NewRequest("GET", "https://yandex.ru", nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Add auth
	req.Header.Add("Authorization", bearer)

	// Create client
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

	// continue work
	fmt.Println(string(body))



}
