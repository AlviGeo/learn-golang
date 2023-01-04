package main

import "fmt"

func main() {
	name := "Alvi"

	switch name {
	case "Alvi":
		fmt.Println("Hello Alvi")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Kenalan donk")
	}

	// switch length := len(name); length > 5 { 
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	// Cara lain yang lebih sederhana, mirip if else
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah panjang")
	}
}