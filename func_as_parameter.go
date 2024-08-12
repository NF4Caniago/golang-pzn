package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("hello", filteredName)
}

func filterMisuh(name string) string {
	if name == "anjing" {
		return "..."
	}else{
		return name
	}
}

func main() {
	name := "anjing"
	sayHelloWithFilter(name, filterMisuh)

	filter := filterMisuh
	sayHelloWithFilter("jokowi", filter)
}