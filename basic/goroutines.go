package main

import (
    "fmt"
    "time"
)

func saludar(nombres []string) {
    for _, nombre := range nombres {
        fmt.Printf("Hola %s\n", nombre)
        time.Sleep(1 * time.Second)
    }
}

func despedir(nombres []string) {
    for _, nombre := range nombres {
        fmt.Printf("Adios %s\n", nombre)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    nombres := []string{"Ana", "Barby", "Carlos", "Dario", "Eli"}
    go saludar(nombres)
    go despedir(nombres)

    var s string
    fmt.Scan(&s)
}
