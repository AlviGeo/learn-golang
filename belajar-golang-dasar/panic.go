package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Alvi")
}

// panic func = func yg bsa digunakan utk menghentikan program(saat err), namun defer tetap akan dieksekusi

// recover func = func yg digunakan utk menangkap data panic
// dgn recover proses panic akan terhenti, sehingga program akan tetap berjalan