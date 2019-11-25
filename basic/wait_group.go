package main

import (
    "fmt"
    "sync"
)

func doSomething(from string, wg *sync.WaitGroup) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
    wg.Done()
}


func main() {
    var wg sync.WaitGroup
    wg.Add(3)

    go doSomething("goroutine1", &wg)

    // funcion anonima 
    go func(msg string) {
        fmt.Println(msg)
        wg.Done()
    }("going")

    go doSomething("goroutine2", &wg)

    wg.Wait()
}
