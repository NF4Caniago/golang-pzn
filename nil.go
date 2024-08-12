package main
import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string {
			"name": "afif",
		}
	}
}

func main() {
	name := NewMap("")
	if name == nil {
		fmt.Println("data masih kosong")
	} else {
		fmt.Println(name)
	}
}