package main

import "fmt"

func main() {
	names := [...]string{"afif", "ilham", "caniago", "latipsa", "lailatur", "eko"}
	slice := names[3:6]
	fmt.Println(slice)

	slice2 := names[:]
	fmt.Println(slice2)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "saptu", "minggu"}
	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "saptu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "libur baru")
	fmt.Println(days)
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)

	var newSlice []string = make([]string , 2, 2)
	newSlice[0] = "afif"
	newSlice[1] = "afif"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "ilham")
	newSlice2 = append(newSlice2, "caniago")
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)
	fmt.Println(days)
	toSlice[4] = "libur coy"
	fmt.Println(toSlice)
	fmt.Println(days)
	fmt.Println(fromSlice)

	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}