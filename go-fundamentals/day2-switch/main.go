package main

import (
	"fmt"
)

/*
📝 EJERCICIO 4: Switch Statement
escribí un programa que:
Parte 1 - Días de la semana:

dayOfWeek := 3
Switch que imprima el día (1=Lunes, 2=Martes, etc.)

Parte 2 - Calificaciones:

grade := 'B' (rune/character)
Switch para convertir letra a puntos:

'A': 4.0
'B': 3.0
'C': 2.0
'D': 1.0
Otro: 0.0
*/

func main() {

	dayOfWeek := 3
	grade := 'B'

	switch dayOfWeek {
	case 1:
		fmt.Println("Lunes.")
	case 2:
		fmt.Println("Martes.")
	case 3:
		fmt.Println("Miercoles.")
	case 4:
		fmt.Println("Jueves.")
	case 5:
		fmt.Println("Viernes.")
	case 6:
		fmt.Println("Sabado.")
	case 7:
		fmt.Println("Domingo.")
	default:
		fmt.Println("Caso no contemplado, error.")
	}

	switch grade {
	case 'A':
		fmt.Printf("Nota %c: %.1f puntos\n", grade, 4.0)
	case 'B':
		fmt.Printf("Nota %c: %.1f puntos\n", grade, 3.0)
	case 'C':
		fmt.Printf("Nota %c: %.1f puntos\n", grade, 2.0)
	case 'D':
		fmt.Printf("Nota %c: %.1f puntos\n", grade, 1.0)
	default:
		fmt.Printf("Nota %c: %.1f puntos\n", grade, 0.0)
	}

}
