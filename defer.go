package main

import "fmt"

func logging() {
	fmt.Println("Apps Selesai")
}

func runApps() {
	defer logging()
	fmt.Println("Apps berjalan")
}

func main() {
	runApps()
}