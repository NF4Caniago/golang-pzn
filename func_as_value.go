package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	vargoodby := getGoodBye
	var2goodbye := getGoodBye
	fmt.Println(vargoodby("afif"))
	fmt.Println(var2goodbye("azizah"))
}