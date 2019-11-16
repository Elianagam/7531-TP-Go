package main

import "fmt"

func main(){
  var mi_var = 100

  var dir_var *int

  dir_var = &mi_var

  fmt.Printf("Valor de la variable 'mi_var': %d \n", mi_var)

  fmt.Printf("Dirección almacenada en 'dir_var': %x \n", dir_var)

  fmt.Printf("Valor de la variable que apunta 'dir_var': %d \n", *dir_var)

  fmt.Printf("Dirección que ocupa el apuntador 'dir_var' en memoria: %x \n", &dir_var)
}
