package main

import (
	"fmt"
	"sync"
	"time"
)

func sayFiveTimes(text string, duration time.Duration, group *sync.WaitGroup){
	for i:=1; i <= 5; i++ {
		fmt.Println(text)
		time.Sleep(duration * time.Millisecond)
	}
	wg.Done()
}

func sayFiveTimesFixed(text string, duration time.Duration, group *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Terminé")

	for i:=0; i <= 5; i++ {
		fmt.Println(text)
		time.Sleep(duration * time.Millisecond)
	}

}


func main() {
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go sayFiveTimesFixed("One", 100, wg)
	go sayFiveTimesFixed("Two", 300, wg)

	wg.Wait()
	//fmt.Println("Terminó!")
}

