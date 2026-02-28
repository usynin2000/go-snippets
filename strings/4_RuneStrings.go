package main

import "fmt"

func main() {
	s := "Привет"
	runes := []rune(s)
	fmt.Println("Длин в символах: ", len(runes)) // 6
	fmt.Println("Первый символ: ", string(runes[0]))
}

// Why to work through runes? What does it mean?

// In Go strings (string) are stored as a sequence of bytes, not characters.
// Each byte is just a number from 0 to 255.

// In Go rune is a synonym for the int32 type, which stores the Unicode code point (character).

// When we do:

// runes := []rune(s)
// Go takes the string in UTF-8 and decodes it into a sequence of characters (rune).
// Then "Привет" becomes an array of 6 rune, where each is one Unicode character (П, р, и and so on).


// When to use rune

// - When working with Unicode characters (russian letters, emojis, chinese characters and so on).

// - When you need to get the correct length of the string in characters.

// - When you need to work with the string element by element: string(runes[i]).

// 👉 If working with pure ASCII (english text, numbers, symbols), then you can use bytes ([]byte).
// But if the string can contain any languages and emojis - it is safer to work through rune.
