package main

import "fmt"

func twice(value int) int {
	return value * 2
}
func apply(value int, operation func(int) int) int {
	return operation(value)
}
func main() {
	operation := twice
	fmt.Println(operation(3), apply(4, operation))
	increment := 1
	next := func(value int) int { return value + increment }
	fmt.Println(next(3))
	increment = 5
	fmt.Println(next(3))
}
