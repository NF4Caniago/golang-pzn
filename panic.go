package main
import "fmt"

func endApps() {
	fmt.Println("End Apps")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApps(myError bool) {
	defer endApps()
	if myError {
		panic("ERRORR!!")
	}
}

func main() {
	runApps(true)
	fmt.Println("afif ilham")
}