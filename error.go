package main
import (
	"fmt"
	"errors"
)

func Pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan noi")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("hasil : ", hasil)
	} else {
		fmt.Println(err.Error())
	}
}