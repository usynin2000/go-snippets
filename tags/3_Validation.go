package main

import (
	"fmt"
	"log"

	"github.com/asaskevich/govalidator"
)

type Example struct {
	Name  string `valid:"-"`
	Email string `valid:"email"`
	URL   string `valid:"url"`
}

func main() {
	// check if email and URL are valid

	result, err := govalidator.ValidateStruct(Example{
		Email: "johndoe@example.com",
		URL:   "https://www.ya.ru",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
