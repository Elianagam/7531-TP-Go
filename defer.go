package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printAndSayGoodbye(text string) {
	defer fmt.Println("Goodbye")
	fmt.Println(text)
}

func sayFiveTimes(text string, seconds int64){
	// Que pasa si el codigo del medio rompe y no llega a ejecutar el Done?
	for i:=1; i <= 5; i++ {
		fmt.Println(text)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}

func sayFiveTimesFixed(text string) {
	defer wg.Done()

	for i:=0; i <= 5; i++ {
		fmt.Println(text)
		time.Sleep(100 * time.Millisecond)
	}

}


func main() {
	printAndSayGoodbye("Hello")

	wg.Add(2)
	go sayFiveTimes("One", 1)
	go sayFiveTimes("Two", 1)
	go sayFiveTimes("Three", 3)

	wg.Wait()
}

