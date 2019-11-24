package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			time.Sleep(time.Second*1)
			x, y = y, x+y
		case <- quit:
			fmt.Println("Salgo")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("Arranqué a leer")
		for i := 0; i < 10; i++ {
			fmt.Println("Leí", <-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
