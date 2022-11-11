package main

import "fmt"

func main() {
	var name string

	name = "Alvi Geovanny"
	fmt.Println(name)

	name = "Alvi G"
	fmt.Println(name)

	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	// tanda : adalah sebagai ganti deklarasi var
	country := "Indonesia"
	fmt.Println(country)

	var (
		firstname = "Alvi"
		lastname = "Geovanny"
	)

	fmt.Println(firstname, lastname)
}