package main
import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var alamat1 Address = Address{"bukittinggi", "Sumbar", "indonesia"}
	var alamat2 *Address = &alamat1

	alamat2.City = "agam"
	fmt.Println(alamat1)
	fmt.Println(alamat2)

	// alamat2 = &Address{"medan", "Sumut", "indonesia"}
	// fmt.Println(alamat1)
	// fmt.Println(alamat2)

	*alamat2 = Address{"jogja", "DIY Yogyakarta", "indonesia"}
	fmt.Println(alamat1)
	fmt.Println(alamat2)
}