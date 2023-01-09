package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Alvi Geovanny", "Alvi"))
	fmt.Println(strings.Contains("Alvi Geovanny", "Budi"))

	fmt.Println(strings.Split("Alvi Geovanny", " "))

	fmt.Println(strings.ToLower("Alvi Geovanny"))
	fmt.Println(strings.ToUpper("Alvi Geovanny"))
	fmt.Println(strings.ToTitle("alvi geovanny"))
	fmt.Println(strings.Trim("   Alvi Geovanny   ", " "))
	fmt.Println(strings.ReplaceAll("Alvi Alvi Alvi", "Alvi", "Joko"))
}