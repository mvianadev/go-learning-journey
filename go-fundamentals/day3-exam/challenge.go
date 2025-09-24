package main

import "fmt"

/*
En challenge.go, implementá un programa que:

Variables iniciales:

studentName := "Juan Pérez"
mathScore := 85
englishScore := 92
scienceScore := 78
attendance := 88 (porcentaje de asistencia)


Lógica a implementar:

Calcular promedio de las 3 materias
Determinar categoría basada en promedio (como Día 2)
PERO: si asistencia < 80%, automáticamente "Reprobado"
Contar cuántas materias están arriba de 80
Usar loops para mostrar cada nota individual
*/

func main() {

	studentName := "Juan Perez"
	mathScore := 85
	englishScore := 92
	scienceScore := 78
	attendance := 88 /*porcentaje*/

	average := (mathScore + englishScore + scienceScore) / 3

	if attendance >= 80 {

		if average >= 90 {
			fmt.Printf("Promedio Excelente: %d\n", average)
		} else if average >= 80 {
			fmt.Printf("Promedio Muy bueno: %d\n", average)
		} else if average >= 70 {
			fmt.Printf("Promedio Bueno: %d\n", average)
		} else if average >= 60 {
			fmt.Printf("Promedio Regular: %d\n", average)
		} else if average < 60 {
			fmt.Printf("Promedio Insuficiente: %d\n", average)
		}

	} else {
		fmt.Println("Reprobado por asistencia menor al 80%.")
	}

	subjects := []string{"Matemáticas", "Inglés", "Ciencias"}
	asignatureOverEight := []int{mathScore, englishScore, scienceScore}
	i := 0
	for index, value := range asignatureOverEight {
		if value > 80 {
			i += 1
		}
		fmt.Printf("Nota: %d materia: %s \n", value, subjects[index])
	}

	fmt.Printf("Asignaturas arriba del 80: %d del estudiante %s", i, studentName)
}
