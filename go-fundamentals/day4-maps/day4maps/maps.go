package day4maps

import (
	"fmt"
)

func BasicMaps() {
	scores := make(map[string]int)

	scores["Ana"] = 95
	scores["Bob"] = 87
	scores["Carlos"] = 92

	fmt.Println("Score de Ana:", scores["Ana"])

	ages := map[string]int{
		"Ana":    26,
		"Bob":    26,
		"Carlos": 26,
	}

	fmt.Println(ages)

	products := map[string]float64{
		"Laptop":     255.99,
		"Televisor":  399.99,
		"microondas": 69.99,
	}

	for k, v := range products {
		fmt.Printf("%s cuesta: $%.2f\n", k, v)
	}
}
