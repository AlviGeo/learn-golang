package main

import "fmt"

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Alvi"
	middleName = "Geo"
	lastName = "Vanny"
	return
}


func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}