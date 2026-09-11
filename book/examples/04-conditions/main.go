package main

import "fmt"

func main() {
	published := true
	for edition := 1; edition <= 3; edition++ {
		if !published {
			fmt.Println("Книга готовится")
			break
		}
		if edition == 2 {
			fmt.Println("Издание", edition, "снято с продажи")
			continue
		}
		fmt.Println("Доступно издание", edition)
	}
}
