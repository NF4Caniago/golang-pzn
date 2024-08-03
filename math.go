package main

import "fmt"

func main() {
	var a = 10
	var b = 2
	var c = a + b
	d := c - b
	e := a * c
	f := a / b
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	a += 4
	b -= 2
	c *= b
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}