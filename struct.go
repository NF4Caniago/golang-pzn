package main

import "fmt"

type Customer struct {
	Name, Address 	string
	Age 			int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("hello", name, "my name is", customer.Name)
}

func main() {
	var afif Customer
	fmt.Println(afif)

	afif.Name = "afif ilham caniago"
	afif.Address = "tebet"
	afif.Age = 27

	fmt.Println(afif)
	fmt.Println(afif.Name)
	fmt.Println(afif.Age)

	tipong := Customer{
		Name: "latipsa",
		Address: "biaro",
		Age: 21,
	}
	fmt.Println(tipong)

	rahmi := Customer{"ramhi", "bandung", 25}
	fmt.Println(rahmi)

	rahmi.sayHello("afif")
}