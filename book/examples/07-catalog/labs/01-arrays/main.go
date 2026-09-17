package main

import "fmt"

func main() {
	formats := [3]string{"PDF", "EPUB", "HTML"}
	other := formats
	other[0] = "TXT"
	fmt.Println(len(formats), formats[0], other[0])
	for index, format := range formats {
		fmt.Println(index, format)
	}
}
