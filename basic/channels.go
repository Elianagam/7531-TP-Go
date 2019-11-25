package main

import (
	"fmt"
	"sync"
	"time"
)

type Comida struct {
	Comida string
	Mesa int
	Dificultad time.Duration
}

func consumer(ch chan Comida, group *sync.WaitGroup) {
	for comidaRecibida := range ch{
		fmt.Printf("La mesa %d recibió %v\n", comidaRecibida.Mesa, comidaRecibida.Comida)
	}
	group.Done()
}

func producer(ch chan Comida, lista_comidas []Comida, group *sync.WaitGroup) {
	for _, comida := range lista_comidas {
		time.Sleep(time.Millisecond * comida.Dificultad)
		ch <- comida
	}
	close(ch) // Qué pasa si no cierro el canal?
	group.Done()
}

func main() {
	comida1 := Comida{"Milanesas con papas fritas", 1, 200}
	comida2 := Comida{"Cazuela de mariscos", 3, 2000}
	comida3 := Comida{"Pizza", 2, 1000}
	lista_comidas := []Comida{comida1, comida2, comida3}

	ch := make(chan Comida, 2)
	//done := make(chan bool, 1)

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go producer(ch, lista_comidas, wg)
	go consumer(ch, wg)
	wg.Wait()
}
