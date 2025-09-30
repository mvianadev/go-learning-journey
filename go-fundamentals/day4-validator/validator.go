package main

import (
	"fmt"
	"strings"
)

type User struct {
	Username string
	Email    string
	Age      int
}

func (u User) Validate() (valid bool, errors []string) {
	// Username: mínimo 3 caracteres
	// Email: debe contener "@"
	// Age: entre 18 y 100
	// Agregar errores al slice si fallan validaciones
	var err []string
	if len(u.Username) < 3 {
		err = append(err, "Username: mínimo 3 caracteres")
	}
	if !strings.Contains(u.Email, "@") {
		err = append(err, "Email: debe contener '@'")
	}
	if u.Age < 18 || u.Age > 100 {
		err = append(err, "Age: entre 18 y 100")
	}

	return len(err) == 0, err
}

func validateUsers(users ...User) []string {
	// Retornar todos los errores de todos los users
	var errores []string
	for _, x := range users {
		_, errs := x.Validate()
		prefijo := fmt.Sprintf("User '%s': ", x.Username)
		for _, err := range errs {
			errores = append(errores, prefijo+err)
		}
	}

	return errores
}

func main() {

	users := []User{
		{Username: "Ana", Email: "ana@example.com", Age: 22},   // ✅ válido
		{Username: "Bo", Email: "boexample.com", Age: 30},      // ❌ email sin @
		{Username: "C", Email: "c@correo.com", Age: 17},        // ❌ username corto + edad inválida
		{Username: "Dani", Email: "dani@correo", Age: 101},     // ❌ email sin dominio + edad inválida
		{Username: "Eli", Email: "eli@correo.com", Age: 18},    // ✅ límite inferior
		{Username: "Fran", Email: "fran@correo.com", Age: 100}, // ✅ límite superior
		{Username: "Gus", Email: "gus@correo.com", Age: 0},     // ❌ edad inválida
		{Username: "Hugo", Email: "hugocorreo.com", Age: 50},   // ❌ email sin @
		{Username: "Isa", Email: "isa@example.com", Age: 99},   // ✅ válido
	}

	for _, u := range users {
		valid, errs := u.Validate()
		fmt.Printf("User '%s' válido: %v\n", u.Username, valid)
		for _, e := range errs {
			fmt.Println("  -", e)
		}
	}

	errores := validateUsers(users...)
	for _, e := range errores {
		fmt.Println(e)
	}

}
