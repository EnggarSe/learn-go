package main

import "fmt"

func main() {
	fmt.Println("---- Dibawah Adalah Tugas Konfersi Dolar -------")
	konfersiDolar()
	fmt.Println("------------------------------------------------")
}

func konfersiDolar() {
	var dolar, rupiah int
	dolar = 15524
	fmt.Print("Masukan Uang Dalam Dolar : ")
	fmt.Scanln(&rupiah)
	uang := rupiah * dolar

	fmt.Printf("%d dolar = %d rupiah \n", rupiah, uang)
}
