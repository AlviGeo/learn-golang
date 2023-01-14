package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// Cara pertama 
	// channel <- "Alvi"
	// data := <- channel
	
	// Cara kedua
	// channel <- "Alvi"
	// fmt.Println(<- channel)
	
	// close(channel)

	go func ()  {
		time.Sleep(2 * time.Second)
		channel <- "Alvi Geovanny"
		fmt.Println("Selesai mengirim Data ke channel")
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

// go test -v -run FuncName