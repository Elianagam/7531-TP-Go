package main

import "fmt"

func main(){

    /* Con Indice */
    fmt.Printf("Imprimo el valor y el indice\n")
    array := []int{1, 2, 3, 4, 5}
    for i, s := range array {
        fmt.Println(i, s)
    }

    /* Solo valor */
    fmt.Printf("\nImprimo solo el valor\n")
    for s:= range array {
        fmt.Printf("-> %d\n",s)
    }

    /* Estilo C */
    fmt.Printf("\nEstilo C\n")
    for i:=0; i < 10; i++ {
        fmt.Printf("-> %d\n", i)
        if i == 5 {
            fmt.Printf("Salimos del ciclo...\n")
            break
        }
    }

    /* While loop */
    i := 1;
    for i < 5 {
        i *= 2
        fmt.Printf("-> %d\n", i)
    }
    fmt.Printf("\n")
}

