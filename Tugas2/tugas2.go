package main

import "fmt"

func main() {
	fmt.Println("---- Dibawah Adalah Tugas Lingkaran ------------")
	luasLingkaran()
	fmt.Println("------------------------------------------------")
}
func luasLingkaran() {
	var jari float32
	fmt.Print("Masukan Jari-jari Lingkaran : ")
	fmt.Scanln(&jari)
	luas := 3.14 * jari * jari
	fmt.Println("Luas Lingkaran : ", luas)
	keliling := 3.14 * 2 * jari
	fmt.Println("Keliling Lingkaran : ", keliling)

}
