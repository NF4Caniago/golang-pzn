package main

import "fmt"

func main() {
	type NoKtp string

	var ktpAfif NoKtp = "111111"
	var contoh string = "22222"
	var contohKtp NoKtp = NoKtp(contoh)

	fmt.Println(ktpAfif)
	fmt.Println(contohKtp)
}