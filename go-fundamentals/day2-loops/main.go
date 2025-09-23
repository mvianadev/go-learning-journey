package main

import (
	"fmt"
)

/*
Ejercicio: Implementá estos 4 tipos de loops:

For clásico: Imprimí números del 1 al 5 ✅
For como while: Cuenta regresiva desde 10 hasta 1 ✅
For infinito con break: Para cuando encuentres el número 7 ✅
Suma acumulativa: Suma números del 1 al 100 ✅
*/

func main() {

	fmt.Println("for loop:")
	for i := 1; i < 6; i++ {
		fmt.Println(i)
	}

	i := 10

	fmt.Println("While loop:")
	for i > 0 {
		fmt.Println(i)
		i -= 1
	}

	fmt.Println("For infinito con break:")

	y := 0
	for { // Loop infinito (sin condición)
		y++
		if y == 7 {
			break // Para cuando encuentre 7
		}
		fmt.Printf("Buscando... %d\n", y)
	}
	fmt.Printf("¡Encontré el 7! Valor: %d\n", y)

	fmt.Println("for loop with sum: ")

	total := 0

	for i := 1; i <= 100; i++ {
		total += i
	}

	fmt.Println(total)
}
