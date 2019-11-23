package main

import (
    "fmt"
)

func saludar(nombres []string) {
    for _, nombre := range nombres {
        fmt.Printf("Hola %s\n", nombre)
    }
}

func despedir(nombres []string) {
    for _, nombre := range nombres {
        fmt.Printf("Adiós %s\n", nombre)
    }
}

func main() {
    nombres := []string{"Orlando", "Daniela", "José", "Carlos", "Andrea", "David", "Carmen"}
    go saludar(nombres)
    go despedir(nombres)
}
