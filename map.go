package main

import "fmt"

func main() {
	person := map[string]string{
		"name" : "afif",
		"address" : "biaro",
	}

	fmt.Println(person)
	fmt.Println(person["name"])

	book := make(map[string]string)
	book["author"] = "afif"
	book["title"] = "jamban keadilan"
	book["salah"] = "pemain bola"

	fmt.Println(book)

	delete(book, "salah")

	fmt.Println(book)

}