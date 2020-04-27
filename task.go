package main

import (
	"fmt"
	"time"
)

/* ini adalah comment */
func main() {
	fmt.Println("---------- Dibawah Adalah Tugas FizzBuzz -------")
	fizzBuzz()
	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("---- Dibawah Adalah Tugas Lingkaran ------------")
	luasLingkaran()
	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("---- Dibawah Adalah Tugas Konfersi Dolar -------")
	konfersiDolar()
	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("---- Dibawah Adalah Tugas Selisih Tanggal -------")
	selisihTanggal()
	fmt.Println("------------------------------------------------")
}
func fizzBuzz() {
	var batas int
	fmt.Print("Masukan Batas Yang Diinginkan : ")
	fmt.Scanln(&batas)

	for i := 1; i <= batas; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
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

func konfersiDolar() {
	var dolar, rupiah int
	dolar = 15524
	fmt.Print("Masukan Uang Dalam Dolar : ")
	fmt.Scanln(&rupiah)
	uang := rupiah * dolar

	fmt.Printf("%d dolar = %d rupiah \n", rupiah, uang)
}

func selisihTanggal() {

	firstDate := time.Date(2019, time.September, 18, 24, 0, 0, 0, time.UTC)
	seconDate := time.Date(2019, time.November, 1, 24, 0, 0, 0, time.UTC)

	diff := seconDate.Sub(firstDate)

	days := int(diff.Hours() / 24)

	fmt.Println("Jarak Antara ", firstDate, " dengan ", seconDate, " adalah = ", days, " hari")
}
