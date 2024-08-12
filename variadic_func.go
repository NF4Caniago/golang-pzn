package main

import "fmt"

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total = number + total
	}

	return total
}

func main() {
	fmt.Println(sum(10,10,10,10))

	numbers := []int{10,10,10,10,10}
	fmt.Println(sum(numbers...))
}