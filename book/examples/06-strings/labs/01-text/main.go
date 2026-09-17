package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	text := "ӘGo"
	var letter rune = 'Ә'
	fmt.Println(len(text), utf8.RuneCountInString(text))
	fmt.Printf("%T %c %d\n", letter, letter, letter)
	fmt.Printf("Первый байт: %d\n", text[0])
	for offset, value := range text {
		fmt.Printf("%d %c\n", offset, value)
	}
	fmt.Printf("%q\n", "Go\nshop")
	fmt.Printf("%q\n", `Go\nshop`)
	old := text
	text = "Go" + " shop"
	fmt.Println(old, text)
}
