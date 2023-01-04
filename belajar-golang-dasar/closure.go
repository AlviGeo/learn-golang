package main

import "fmt"

func main() {
	name := "Budi"
	counter := 0

	increment := func() {
		// jika menggunakan name = "Alvi", maka var name diatas juga ikut diganti
		// jadi harus menggunakan :=
		name := "Alvi"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}