package main

import (
	"fmt"
	"math"
)

type Figura interface {
	GetArea() float64
}

type Circulo struct {
	Radio float64
}

func (c *Circulo) GetArea() float64 {
	return math.Pi * math.Pow(c.Radio, 2)
}

type Cuadrado struct {
	Lado float64
}

func (c *Cuadrado) GetArea() float64 {
	return math.Pow(c.Lado, 2)
}

func main() {
	var circulo Figura = &Circulo{Radio: 5}
	var cuadrado Figura = &Cuadrado{Lado: 5}

	fmt.Printf("El área del circulo es: %f \n", circulo.GetArea())
	fmt.Printf("El área del cuadrado es: %f \n", cuadrado.GetArea())
}