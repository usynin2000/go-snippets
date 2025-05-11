package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


type RequestData struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type ResponseData struct {
	Message string `json:message"`
}

func main() {
	reqData := RequestData{
		Name: "Alice",
		Age: 30,
	}

	jsonBody, err := json.Marshal(reqData)

	if err != nil {
		panic(err)
	}

	url := "https://httpbin.org/post" // echo-сервер, который возвращает нам наши же данные

	req, err := http.NewRequest("POST",  url, bytes.NewBuffer(jsonBody))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("== Сарой ответ ==")
	fmt.Println(string(body))

	type HttpBinResponse struct {
		JSON RequestData `json:"json"`
	}

	var parsed HttpBinResponse
	err = json.Unmarshal(body, &parsed)

	if err != nil {
		panic(err)
	}

	fmt.Println("\n== Распарсенные данные ==")
	fmt.Printf("Name: %s, Age: %d\n", parsed.JSON.Name, parsed.JSON.Age)
}