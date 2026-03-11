package main

import (
	"log"
	"os"
)

func main() {
	// create file info.log and handle error if something went wrong
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// defer closing file
	defer file.Close()

	// set output destination to file info.log
	log.SetOutput(file)
	log.Print("Logging to file in GO!")


}
