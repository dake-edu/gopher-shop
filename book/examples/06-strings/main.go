package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	title, ok := normalizeTitle("  Шаңырақ Go  ")
	if !ok {
		fmt.Println("Название не подходит")
		return
	}
	fmt.Printf("Название: %q\n", title)
	fmt.Println("Байтов:", len(title))
	fmt.Println("Рун:", utf8.RuneCountInString(title))
	for offset, letter := range "ӘGo" {
		fmt.Printf("Байт %d: %c\n", offset, letter)
	}
}
