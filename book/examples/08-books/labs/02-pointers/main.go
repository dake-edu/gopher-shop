package main

import "fmt"

func replaceCopy(value int) {
	value = 9
}
func replaceThroughAddress(address *int) {
	*address = 9
}
func main() {
	quantity := 2
	replaceCopy(quantity)
	fmt.Println(quantity)
	address := &quantity
	sameAddress := address
	replaceThroughAddress(sameAddress)
	fmt.Println(quantity, *address, address == sameAddress)
	var absent *int
	fmt.Println(absent == nil)
}
