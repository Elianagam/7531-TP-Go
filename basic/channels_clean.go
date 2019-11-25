package main

import (
	"fmt"
	"time"
)

type Comida struct {
	Comida string
	Mesa int
	Dificultad time.Duration
}

func consumer(ch <-chan Comida, done chan bool) {
	for i := 3; i < 3; i++ {
		comida := <-ch
		fmt.Printf("La mesa %d recibió %v\n", comida.Mesa, comida.Comida)
	}
	done <- true
}

func producer(ch chan<- Comida, lista_comidas []Comida) {
	for _, comida := range lista_comidas {
		time.Sleep(time.Millisecond * comida.Dificultad)
		ch <- comida
	}
	close(ch) // Qué pasa si no cierro el canal?
}

func main() {
	comida1 := Comida{"Milanesas con papas fritas", 1, 200}
	comida2 := Comida{"Cazuela de mariscos", 3, 1500}
	comida3 := Comida{"Pizza", 2, 1000}
	lista_comidas := []Comida{comida1, comida2, comida3}

	ch := make(chan Comida, 3)
	done := make(chan bool)

	go producer(ch, lista_comidas)
	go consumer(ch, done)

	<-done
}
