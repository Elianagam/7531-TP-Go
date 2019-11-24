package main

import (
	"fmt"
	"time"
)

func saveNumbers(ch chan int, done chan bool) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 2)
		ch <- i
	}
	done <- true
}

func main() {

	ch := make(chan int, 2)
	done := make(chan bool, 1)
	go saveNumbers(ch, done)

	fmt.Println("Estoy esperando arrancar")

	readChannel:
		for {
			select {
			case number := <-ch:
				fmt.Println(number)
			case <-done:
				fmt.Println("Terminó!")
				break readChannel
			}
		}
}
