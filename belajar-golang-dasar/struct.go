package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func main() {
	var eko Customer
	eko.Name = "Alvi"
	eko.Address = "Indonesia"
	eko.Age = 20

	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	joko := Customer {
		Name: "Joko",
		Address: "Cirebon",
		Age: 35,
	}

	fmt.Println(joko)

	// cara ini not recommended karena tergantung dengan posisinya struct
	budi := Customer{"Budi", "Jakarta", 35}
	fmt.Println(budi)
}