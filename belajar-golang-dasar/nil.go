package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string {
			"name": name,
		}
	}
}

func main() {
	// var person map[string]string = nil
	// person := NewMap("Alvi")
	// fmt.Println(person)

	// Cara pertama
	// var person map[string]string
	// if person["name"] == "" {
	// 	fmt.Println("Data kosong")
	// } else {
	// 	fmt.Println(person)
	// }

	// Cara kedua (lebih mudah)
	var person map[string]string = NewMap("Alvi")
	if person == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person)
	}
}

// Nil hanya bisa digunakan pada interface, func, map, slice, pointer, dan channel