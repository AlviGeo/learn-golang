package main

import "fmt"

type Address struct {
	City, Province, Country string 
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	// Function new
	var address4 *Address = new(Address)
	address4.City = "Pekanbaru"
	fmt.Println(address4)

	var alamat = Address {
		City: "Pekanbaru",
		Province: "Riau",
		Country: "",
	}

	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}