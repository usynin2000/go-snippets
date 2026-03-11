package main

import (
	"log"
	"os"
)

func main() {
	// create file info.log and handle error if something went wrong
	file, err := os.OpenFile(
		"info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644,
	)
	if err != nil {
		log.Fatal(err)
	}

	// defer closing file

	defer file.Close()

	log.SetOutput(file)

	 // add flags for output date, time, file name
	 log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	 log.Print("loggin in file GO")






}
