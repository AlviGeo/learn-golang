package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

// Dua bawah ini merupakan struct function
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuu from", a.Name)
}

func main() {
	var eko Customer
	eko.Name = "Alvi"
	eko.Address = "Indonesia"
	eko.Age = 20

	eko.sayHi("Joko")
	eko.sayHuuu()

	// fmt.Println(eko.Name)
	// fmt.Println(eko.Address)
	// fmt.Println(eko.Age)

	// joko := Customer {
	// 	Name: "Joko",
	// 	Address: "Cirebon",
	// 	Age: 35,
	// }

	// fmt.Println(joko)

	// // cara ini not recommended karena tergantung dengan posisinya struct
	// budi := Customer{"Budi", "Jakarta", 35}
	// fmt.Println(budi)
}