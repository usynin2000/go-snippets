package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

func main() {
	client := resty.New()


	client.
		// устанавливаем количестов повторений
		SetRetryCount(3).
		// длительность ожидания между попытками
		SetRetryWaitTime(30 * time.Second).
		// длительность максимального ожидания
		SetRetryMaxWaitTime(90 * time.Second)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"title": "foo", "body": "bar", "userId": 7}`).
		Post("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)


	// другой вариант POST запроса
	// если передается map, то по умолчанию используется json
	resp, err = client.R().
		SetBody(map[string]interface{}{"title":"My title", "body":"Content", "userId": 7}).
		Post("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)


}