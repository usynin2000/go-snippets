package main

import (
	"flag"
	"fmt"
)

func main() {
	// указываем имя флага, значение по умолчанию и описание
	// specify flag name, default value and description
	imgFile := flag.String("file", "", "input image file")

	// parse command line arguments
	flag.Parse()
	fmt.Println("Image file:", *imgFile)
}

// go run 2_FileFlag.go -file=some_image.png  
// go run 2_FileFlag.go --file=some_image.png
