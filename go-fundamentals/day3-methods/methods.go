package main

import (
	"fmt"
)

func main() {

	fmt.Println("Calculando")
	rect := Rectangle{Width: 10.5, Height: 20.3}
	fmt.Println("Area: ", rect.Area())
	fmt.Println("Perimetro: ", rect.Perimeter())
	fmt.Println(rect.Description())

	fmt.Println("Antes:", rect)

	rect.Scale(2)                 // Duplicar tamaño
	fmt.Println("Después:", rect) // ¿Cambió?

}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Description() string {
	return fmt.Sprintf("Un rectangulo de ancho: %.2f y de alto: %.2f", r.Width, r.Height)
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}
