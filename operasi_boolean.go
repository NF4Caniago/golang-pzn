package main
import "fmt"
func main() {
	var nilaiAbsensi = 82
	var nilaiAkhir = 90
	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = nilaiAbsensi > 80
	lulus := lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)
}