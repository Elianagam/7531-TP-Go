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

func sayFiveTimes(text string){
	// Que pasa si el código del medio rompe y no llega a ejecutar el Done?
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
	//printAndSayGoodbye("Hello")

	wg.Add(3)
	go sayFiveTimes("One", 1)
	go sayFiveTimes("Two", 1)

	wg.Wait()
	fmt.Println("Terminó!")
}

