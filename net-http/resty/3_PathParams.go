package main

import (
	"fmt"
	"reflect"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()

	resp, err := client.R().SetPathParams(map[string]string{
		"postID": "1",
	}).Get("https://jsonplaceholder.typicode.com/posts/{postID}")

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	fmt.Println(reflect.TypeOf(resp))
}
