package main

import "fmt"

type Person struct {
	FirstName, LastName string
	Age       int
}

func main() {
	//var p = Person{"Eliana", "Gamarra", 22}

	//var x = Person{"Eliana"} /* Error, no Compila */
	//var y = Person{FirstName: "Eliana", LastName: "Gamarra", Age: 22}

	var z = Person{FirstName: "Eliana"} // LastName: "", Age: 0

	fmt.Println(z)
}

