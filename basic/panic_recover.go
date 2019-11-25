package main

import(
	"fmt"
	"math"
	"sync"
)

var wg sync.WaitGroup

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Agarré el error: ", r)
	}

	fmt.Println("Sigo el código..")
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func cantBePrime() {
	defer wg.Done()
	defer catchErr()

	for i := 1; i <= 5; i++ {
		if isPrime(i) {
			panic("Es un número primo!")
		}
		fmt.Println(i)
	}
}


func main() {
	wg.Add(1)
	go cantBePrime()
	wg.Wait()
}
