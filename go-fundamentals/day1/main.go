package main

import (
	"fmt"
)

func main() {
	age := 25
	name := "Mauricio"
	salary := 50000.50
	isEmployed := true

	fmt.Printf("Hola soy %s, tengo %d y soy empleado %t y cobro %.2f", name, age, isEmployed, salary)
}
