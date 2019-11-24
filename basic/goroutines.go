package main

import (
    "fmt"
   // "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {
    f("direct")

    go f("goroutine1")  

    // funcion anonima 
    go func(msg string) {
        fmt.Println(msg)
    }("going")

    go f("goroutine2") 

    //time.Sleep(time.Second)
    fmt.Println("done")
}





