package main

import (
    "fmt"
    "sync"
)

func f(from string, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}


func main() {
    var wg sync.WaitGroup
    wg.Add(3)

    go f("goroutine1", &wg)  

    // funcion anonima 
    go func(msg string) {
        defer wg.Done()
        fmt.Println(msg)
    }("going")

    go f("goroutine2", &wg) 

    wg.Wait()
}