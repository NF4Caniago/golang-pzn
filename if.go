package main

import "fmt"

func main() {
	name := "afif ilham"
	if name == "afif" {
		fmt.Println(name)
	} else if name == "ilham" {
		fmt.Println("Halo afif")
	} else {
		fmt.Println("Tolong namo")
	}

	if lenght := len(name); lenght > 5 {
		fmt.Println("nama kepanjangan")
	} else {
		fmt.Println("nama kependekan")
	}
}