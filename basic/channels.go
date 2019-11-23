package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string, 1)

	go func() {
		text := LongRunningProcess()
		ch <- text
	}()

	select {
	case res := <- ch:
		fmt.Println(res)
	case <- time.After(2 * time.Second):
		fmt.Println("Me quedé sin tiempo :c")
	}

}

func LongRunningProcess() string {
	time.Sleep(1 * time.Second)
	return "Terminé de ejecutar!"
}
