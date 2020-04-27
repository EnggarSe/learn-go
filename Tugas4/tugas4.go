package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---- Dibawah Adalah Tugas Selisih Tanggal -------")
	selisihTanggal()
	fmt.Println("------------------------------------------------")
}
func selisihTanggal() {

	firstDate := time.Date(2019, time.September, 18, 24, 0, 0, 0, time.UTC)
	secondDate := time.Date(2019, time.November, 1, 24, 0, 0, 0, time.UTC)

	diff := secondDate.Sub(firstDate)

	days := int(diff.Hours() / 24)

	fmt.Println("Jarak Antara ", firstDate, " dengan ", secondDate, " adalah = ", days, " hari")
}
