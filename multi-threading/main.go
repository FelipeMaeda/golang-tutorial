package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int)

	// Function Assincrona
	go func() {
		for i := 1; i <= 50; i++ {
			go worker(channel)
		}
	}()
	// Anonymous Function

	for i := 0; i <= 100; i++ {
		channel <- i
	}

}

func worker(channel chan int) {
	for i := range channel {
		fmt.Println(i)
		time.Sleep(time.Second * 2)
	}
}
