package day4maps

import "fmt"

type Student struct {
	Name  string
	Grade float64
	Age   int
}

func StructsMaps() {

	// Map donde la clave es ID (string) y el valor es Student
	students := make(map[string]Student)

	// TODO: Agregar 3 estudiantes
	students["S001"] = Student{Name: "Andres", Grade: 7.99, Age: 21}
	students["S002"] = Student{Name: "Valeria", Grade: 5.59, Age: 19}
	students["S003"] = Student{Name: "Mauricio", Grade: 8.99, Age: 24}

	// TODO: Buscar un estudiante por ID
	fmt.Println(students["S002"])
	// TODO: Actualizar la nota de un estudiante
	students["S002"] = Student{
		Name:  students["S002"].Name,
		Age:   students["S002"].Age,
		Grade: 7.25,
	}
	// TODO: Calcular promedio de todas las notas
	var i int
	var total float64
	for _, v := range students {
		i++
		total += v.Grade
	}
	var promedio float64
	if i > 0 {
		promedio = total / float64(i)
		fmt.Printf("Promedio de %d estudiantes: %.2f\n", i, promedio)
	} else {
		fmt.Println("No hay estudiantes válidos para calcular el promedio.")
	}
	// TODO: Eliminar un estudiante
	delete(students, "S002")
	fmt.Println(students)
}
