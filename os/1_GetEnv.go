package main

import (
	"fmt"
	"os"
)

func main() {
	u := os.Getenv("USER")

	key, is_env := os.LookupEnv("SHELL")


	fmt.Println(u)
	fmt.Println(key, is_env)
}