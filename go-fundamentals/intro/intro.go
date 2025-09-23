package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("🚀 ¡Bienvenido a Go!")

	// Variables - Go infiere el tipo
	name := "QA Developer"
	experience := 5
	isLearningGo := true

	// Printf formateado
	fmt.Printf("Soy %s con %d años de experiencia\n", name, experience)
	fmt.Printf("¿Estoy aprendiendo Go? %t\n", isLearningGo)

	// Llamar funciones
	greeting := createGreeting("Go Developer")
	fmt.Println(greeting)

	// Goroutines - concurrencia básica
	fmt.Println("\n🔄 Probando concurrencia:")
	go printWithDelay("Goroutine 1", 1)
	go printWithDelay("Goroutine 2", 2)

	// Esperar para ver las goroutines
	time.Sleep(3 * time.Second)
	fmt.Println("✅ ¡Primer programa completado!")
}

// Función que retorna un valor
func createGreeting(role string) string {
	return fmt.Sprintf("¡Hola futuro %s!", role)
}

// Función que simula trabajo concurrente
func printWithDelay(message string, seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Printf("⚡ %s ejecutada después de %d segundos\n", message, seconds)
}
