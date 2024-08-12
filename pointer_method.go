package main
import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	var afif Man = Man{"afif"}
	afif.Married()
	fmt.Println(afif.Name)
}