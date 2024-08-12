package main

import "fmt"

func factorialRecusive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecusive(value-1)
	}
}

func main() {
	fmt.Println(factorialRecusive(10))
}