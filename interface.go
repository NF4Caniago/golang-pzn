package main
import "fmt"

type HasName interface {
	GetName() string
}

type Person struct{
	Name string
}

type Animal struct{
	Race, Color string
}

func sayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Race
}

func main() {
	people := Person{"afif"}
	sayHello(people)

	cat := Animal{
		Race: "cat",
		Color: "Black",
	}
	sayHello(cat)
}