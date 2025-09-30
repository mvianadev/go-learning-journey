package main

import (
	"fmt"
	"strings"
)

func countWords(texts ...string) int {
	// Contar total de palabras en todos los textos
	// Hint: strings.Fields(text) divide por espacios
	total := 0

	for _, text := range texts {
		palabras := strings.Fields(text)
		total += len(palabras)
	}
	return total
}

func createLengthFilter(minLength int) func(string) bool {
	// Retorna función que verifica si texto >= minLength
	return func(text string) bool {
		return len(text) >= minLength
	}
}

func processText(text string, filter func(string) bool) {
	defer fmt.Println("Procesamiento completado")
	// Aplicar filter al texto y mostrar resultado
	if filter(text) {
		fmt.Println("Texto válido:", text)
	} else {
		fmt.Println("Texto demasiado corto:", text)
	}
}

func main() {
	var texto1 = "Hola mundo"
	var texto2 = "Esto es QA"
	var texto3 = "MrZadman optimiza builds"

	var textos = []string{texto1, texto2, texto3}

	fmt.Println(countWords(textos...))
	// Salida: 8

	filtro := createLengthFilter(10)
	for _, t := range textos {
		processText(t, filtro)
	}
}
