package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "afif"
	names[1] = "ilham"
	names[2] = "caniago"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [4]int {
		90,
		80,
		70,
	}
	fmt.Println(values)

	var values2 = [4]int {90,80,70}
	fmt.Println(values2)

	var values3 = [...]int {
		1,
		2,
		3,
		4,
	}
	fmt.Println(len(values3))
}