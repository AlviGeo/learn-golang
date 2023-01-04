package main

import "fmt"

func main() {
	var name = "Alvi"

	if name == "Alvi" {
		fmt.Println("Hello Alvi")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, kenalan donk")
	}

	// var length = len(name)
	// if length > 5 {
	// 	fmt.Println("Terlalu panjang")
	// } else {
	// 	fmt.Println("Nama sudah benar")
	// }

	// Cara kedua dari yang atas
	if length := len(name) ;length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}