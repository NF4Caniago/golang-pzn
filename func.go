package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getGetHello(name string) string {
	return "hello " + name
}

func getFullName() (string, string) {
	return "afifi", "ilham"
}

func getFullName2() (firstName, middleName, lastName string) {
	firstName = "afif"
	middleName = "ilham"
	lastName = "caniago"

	return 
}

func main() {
	sayHello("afif", "ilham")
	result := getGetHello("afif")
	fmt.Println(result)
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	myFirstName, _ := getFullName()
	fmt.Println(myFirstName)
	a, b, c := getFullName2()
	fmt.Println(a, b, c)
}