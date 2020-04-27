package main

import "fmt"

func main() {
	fmt.Println("---------- Dibawah Adalah Tugas FizzBuzz -------")
	fizzBuzz()
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
