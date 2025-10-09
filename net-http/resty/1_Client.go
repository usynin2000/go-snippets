package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func main() {
	// create a client

	client := resty.New()

	resp, err := client.R().
		SetAuthToken("Bearer <Token>").
		Get("https://yandex.ru")

	fmt.Println("Исследуем объект Response:")
	fmt.Println("Error    :", err)
	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Status :", resp.Status())
	fmt.Println("Time   :", resp.Time())
	fmt.Println("Recieved At:", resp.ReceivedAt())
	fmt.Println("Body       :\n", resp)
    fmt.Println("----")
}