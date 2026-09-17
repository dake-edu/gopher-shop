package main

import "fmt"

func main() {
	for edition := 1; edition <= 3; edition++ {
		fmt.Println("Издание", edition)
	}
	remaining := 2
	for remaining > 0 {
		fmt.Println("Осталось", remaining)
		remaining = remaining - 1
	}
	number := 0
	for {
		number++
		if number == 2 {
			continue
		}
		if number == 4 {
			break
		}
		fmt.Println("Выбрано", number)
	}
}
