package main

import "fmt"

func main() {
	published := true
	quantity := 0
	if !published {
		fmt.Println("Черновик")
	} else if quantity == 0 {
		fmt.Println("Нет экземпляров")
	} else {
		fmt.Println("Можно выбрать")
	}
	fmt.Println(published && quantity > 0)
	fmt.Println(published || quantity > 0)
	status := "draft"
	switch status {
	case "draft":
		fmt.Println("Готовится")
	case "published":
		fmt.Println("На витрине")
	default:
		fmt.Println("Неизвестное состояние")
	}
}
