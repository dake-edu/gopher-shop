package main

import "fmt"

func observe() {
	value := 1
	defer fmt.Println("argument", value)
	defer func() { fmt.Println("closure", value) }()
	value = 2
	fmt.Println("body", value)
}
func result() (value int) {
	defer func() { value = value + 1 }()
	return 10
}
func main() {
	observe()
	fmt.Println("result", result())
}
