package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)

	// fmt.Println(args[1])
	// fmt.Println(args[2])
	// fmt.Println(args[3])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	
	// Jika ingin set bisa dengan
	// export APP_USERNAME = root
	// export APP_PASSWORD = root

	fmt.Println(username)
	fmt.Println(password)
}

// hasil output dari os.Args kira" seperti ini :
// [/tmp/go-build3141891104/b001/exe/os]

// func lebih banyak bisa lihat di : https://pkg.go.dev/os