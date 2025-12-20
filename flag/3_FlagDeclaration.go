package main

import (
	"fmt"
	"flag"
)


func main() {

	imgFile :=  flag.String("file", "", "input image file")
	destDir := flag.String("dest", "./output", "destination folder")
	width := flag.Int("w", 1024, "width of the image")
	isThumb := flag.Bool("thumb", false, "create thumb")


	// parse command line arguments
	flag.Parse()

	fmt.Println("Image file:", *imgFile)
	fmt.Println("Destination folder:", *destDir)
	fmt.Println("Width:", *width)
	fmt.Println("Thumbs:", *isThumb)
}

// go run 3_FlagDeclaration.go -thumb -w 1024 -file "./in/img0001.jpg" -dest "/home/user/images"

// Image file: ./in/img0001.jpg
// Destination folder: /home/user/images
// Width: 1024
// Thumbs: true