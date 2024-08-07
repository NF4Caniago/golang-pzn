package main

import "fmt"

func main() {
	counter := 10
	for counter > 0 {
		fmt.Println("Perulangan ke", counter)
		counter--
	}

	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println(counter2)
	}

	names := []string{"afif", "ilham", "caniago"}
	for i := 0; i < 3; i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println(name, " index : ", index)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}