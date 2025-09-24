package main

import (
	"fmt"
)

func main() {

	greetUser("Mauricio")

	area := calculateArea(14.24, 5.23)
	fmt.Printf("El area es: %.2f\n", area)

	operation, err := divide(4, 0)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("El resultado de la division:", operation)
	}

	puntaje := [5]int{76, 87, 98, 86, 90}
	promedio := processGrades(puntaje)

	fmt.Println("El promedio es: ", promedio)

	scores := []int{45, 89, 67, 23, 91, 56, 78, -23}
	maxScore := findMax(scores)
	fmt.Printf("El puntaje más alto es: %d\n", maxScore)

	myGrades := []int{85, 90, 78}
	fmt.Println("Notas originales:", myGrades)

	myGrades = addGrade(myGrades, 95)
	fmt.Println("Después de agregar 95:", myGrades)

}

func greetUser(name string) {
	fmt.Printf("Saludos %s.\n", name)
}

func calculateArea(width, height float64) float64 {
	return width * height
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("no se puede divir por cero")
	}

	return a / b, nil
}

func processGrades(grades [5]int) float64 {
	// Calcular promedio de array fijo de 5 elementos
	var promedio float64
	for _, v := range grades {
		promedio += float64(v)
	}

	return promedio / float64(len(grades))
}

func findMax(numbers []int) int {
	// Encontrar el número máximo en slice
	max := numbers[0]

	for _, v := range numbers {
		if v > max {
			max = v
		}
	}

	return max
}

func addGrade(grades []int, newGrade int) []int {
	// Agregar nueva nota al slice
	return append(grades, newGrade)
}
