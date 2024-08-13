package main
import "fmt"

func random() any {
	return "ok"
}

func main() {
	var value any = random()
	var valueString string = value.(string)
	// valueString, status := value.(string)
	// fmt.Println(status)
	// fmt.Println(valueString)

	var result any = random()
	switch typeResult := result.(type) {
	case string:
		fmt.Println("string", typeResult)
	case int:
		fmt.Println("int", typeResult)
	default:
		fmt.Println("Unknown")
	}
}