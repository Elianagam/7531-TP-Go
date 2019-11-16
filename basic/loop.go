package main

import "fmt"

func main(){
  arreglo:=[7]int{0,1,4,6,10,9}

  /* Con Indice */
  for i, j:= range arreglo{
    fmt.Printf("Valor de j: %d en vuelta #%d\n", j,i)
  }

  /* Solo valor */
  for i:= range arreglo{
    fmt.Printf("Valor de i: %d\n", i)
  }

  for i:=0 ; i < 10; i++ {
    fmt.Printf("Valor de i: %d", i)
    if i == 7 {
      fmt.Printf(" así que saldremos del ciclo...\n")
      break
    }
  fmt.Printf("\n")
  }

  CICLO: for i < 10 {
    if i == 6 {
      i = i + 3
      fmt.Println("Saltando a etiqueta CICLO con i = i + 3")
      goto CICLO
    }
    fmt.Printf("Valor de i: %d\n", i)
    i++
  }
}