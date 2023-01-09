package main

import (
	"belajar-golang-dasar/database"
	"fmt"
)

// Jika ingin panggil namun tidak mau eksekusi, tambahkan _
// import _ "belajar-golang-dasar/database"

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}

