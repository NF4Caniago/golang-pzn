package helper
import "fmt"

var version = "1.0.0"
var Application = "Apps Go"

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Halo " + name
}

func Contoh() {
	fmt.Println(sayGoodBye("afif"))
	fmt.Println(version)
}