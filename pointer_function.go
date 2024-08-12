package main
import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCityToIndonesia(address *Address) {
	address.City = "Bukittinggi"
}

func main() {
	var alamat Address
	ChangeCityToIndonesia(&alamat)
	fmt.Println(alamat)
}