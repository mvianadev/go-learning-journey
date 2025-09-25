//EJERCICIO AVANZADO 1: Named Returns
//Implementá esta función usando named returns:

package main

import (
	"fmt"
)

func main() {

	q, r, err := divmod(4, 8)

	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		fmt.Printf("quotient: %d, remainder: %d.", q, r)
	}

	fmt.Println(sum(1, 2, 3))        // 6
	fmt.Println(sum(10, 20, 30, 40)) // 100
	fmt.Println(sum())               // 0 (sin argumentos)

	double := createMultiplier(2)
	triple := createMultiplier(3)

	fmt.Println(double(5)) // 10
	fmt.Println(triple(4)) // 12

	processFile()
	multipleDefers()
}

func divmod(a, b int) (quotient, remainder int, err error) {
	// Si b == 0, retornar error
	// Sino, calcular división y resto
	// Las variables quotient, remainder, err ya están declaradas
	// Solo asignás valores y usás "return" vacío
	if b == 0 {
		err = fmt.Errorf("division por cero")
		return
	}
	quotient = a / b
	remainder = a % b
	return
}

func sum(numbers ...int) int {
	// Sumar todos los números que reciba
	// numbers es un slice []int internamente
	var total int
	for _, v := range numbers {
		total += v
	}

	return total
}

func createMultiplier(factor int) func(int) int {
	// Retorna una función anónima que multiplica por factor
	return func(x int) int {
		return x * factor // "Captura" la variable factor (closure)
	}
}

func processFile() {
	fmt.Println("1. Abriendo archivo")
	defer fmt.Println("4. Cerrando archivo") // Se ejecuta AL FINAL

	fmt.Println("2. Leyendo datos")
	fmt.Println("3. Procesando datos")
	// defer se ejecuta aquí automáticamente
}

func multipleDefers() {
	defer fmt.Println("Primero")
	defer fmt.Println("Segundo")
	defer fmt.Println("Tercero")
	fmt.Println("Normal")
}
