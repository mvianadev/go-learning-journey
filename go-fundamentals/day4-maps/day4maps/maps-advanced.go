package day4maps

import "fmt"

func AdvanceMaps() {
	inventory := map[string]int{
		"manzanas": 50,
		"peras":    30,
		"naranjas": 45,
	}

	// 1. Verificar si una clave existe (importante!)
	// value, exists := inventory["bananas"]
	value, exists := inventory["bananas"]

	if exists {
		fmt.Println("Cantidad de bananas: ", value)
	} else {
		fmt.Println("No hay bananas.")
	}
	fmt.Println(inventory)
	// 2. Eliminar un elemento
	// delete(inventory, "peras")
	delete(inventory, "naranjas")
	fmt.Println(inventory)
	// 3. Obtener tamaño del map
	// len(inventory)
	println(len(inventory))

	// 4. Iterar solo claves O solo valores
	for k := range inventory {
		fmt.Println(k)
	}
}
