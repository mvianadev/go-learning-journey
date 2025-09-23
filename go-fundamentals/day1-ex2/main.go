package main

import (
	"fmt"
)

func main() {
	num1 := 10
	num2 := 3

	suma := num1 + num2
	resta := num1 - num2
	multiplicacion := num1 * num2
	division := num1 / num2
	resto := num1 % num2

	fmt.Printf("La suma es: %d\nLa resta es: %d\nLa multiplicacion es: %d\nLa division es: %d\nEl resto es: %d", suma, resta, multiplicacion, division, resto)

	price := 100
	discount := 0.15

	finalPrice := float64(price) * (1 - discount)
	fmt.Printf("\n\nEl precio con descuento es: %f\n", finalPrice)

	fmt.Printf("%d es menor que %d ", num1, num2)
	fmt.Println(num1 < num2)
	fmt.Printf("%d es mayor que %d ", num1, num2)
	fmt.Println(num1 > num2)
	fmt.Printf("%d es menor igual que %d ", num1, num2)
	fmt.Println(num1 <= num2)
	fmt.Printf("%d es mayor igual que %d ", num1, num2)
	fmt.Println(num1 >= num2)
	fmt.Printf("%d es igual que %d ", num1, num2)
	fmt.Println(num1 == num2)
	fmt.Printf("%d es diferente que %d ", num1, num2)
	fmt.Println(num1 != num2)
}
