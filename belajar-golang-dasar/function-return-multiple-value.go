package main

import "fmt"

func getFullName() (string, string) {
	return "Alvi", "Geovanny"
}

// Jika ingin menghiraukan valuesnya, dapat diganti dengan _ 
func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
