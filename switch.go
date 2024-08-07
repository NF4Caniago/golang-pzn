package main

import "fmt"

func main() {
	name := "afif"

	switch name {
	case "afif":
		fmt.Println(name)
	case "ilham":
		fmt.Println("ini ilham")
	case "caniago":
		
		fmt.Println("ini caniago")
	}

	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama terlalu pendek")
	}

	lenght := len(name)
	switch {
	case lenght > 10:
		fmt.Println("Nama terlalu panjang")
	case lenght > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("benar")
	}
}