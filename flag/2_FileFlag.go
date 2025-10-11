package main

import (
	"flag"
	"fmt"
)

func main() {
	// указываем имя флага, значение по умолчанию и описание
	imgFile := flag.String("file", "", "input image file")

	// делаем разбор командной строки
	flag.Parse()
	fmt.Println("Image file:", *imgFile)
}

// go run 2_FileFlag.go -file=some_image.png  
// go run 2_FileFlag.go --file=some_image.png


// А вот если передать ошибочно:
// s.usynin•testing/go-snippets/flag(master⚡)» go run 2_FileFlag.go -f=some_image.png                             [15:01:13]
// flag provided but not defined: -f
// Usage of /var/folders/6b/w1gc0jz114ldg1lr87jm3qnrsgh5xh/T/go-build3602375422/b001/exe/2_FileFlag:
//   -file string
//         input image file
// exit status 2