package main

import (
	"fmt"
)

/*
Usando if/else if/else, determina la calificación:

90-100: "Excelente"
80-89: "Muy bueno"
70-79: "Bueno"
60-69: "Regular"
Menos de 60: "Insuficiente"


Declara age := 17
Verifica si puede votar (>= 18) y muestra mensaje apropiado
*/

func main() {

	score := 75
	age := 17

	if score >= 90 {
		fmt.Println("Excelente.")
	} else if score >= 80 {
		fmt.Println("Muy bueno.")
	} else if score >= 70 {
		fmt.Println("Bueno.")
	} else if score >= 60 {
		fmt.Println("Regular.")
	} else if score < 60 {
		fmt.Println("Insuficiente.")
	}

	if age >= 18 {
		fmt.Println("Puede votar")
	} else {
		fmt.Println("Usted es menor de edad y no puede votar.")
	}

}
