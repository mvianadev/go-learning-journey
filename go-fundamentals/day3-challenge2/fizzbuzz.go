package main

import "fmt"

/*
FizzBuzz Empresarial (1-100):

Múltiplos de 3: "Fizz"
Múltiplos de 5: "Buzz"
Múltiplos de 7: "Bang"
Combinaciones:

3 Y 5: "FizzBuzz"
3 Y 7: "FizzBang"
5 Y 7: "BuzzBang"
3, 5 Y 7: "FizzBuzzBang"


Otros números: el número mismo
*/

func main() {

	for i := 1; i < 150; i++ {

		if i%3 == 0 && i%5 == 0 && i%7 == 0 {
			fmt.Printf("%d: FizzBuzzBang\n", i)
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%d: FizzBuzz\n", i)
		} else if i%3 == 0 && i%7 == 0 {
			fmt.Printf("%d: FizzBang\n", i)
		} else if i%5 == 0 && i%7 == 0 {
			fmt.Printf("%d: BuzzBang\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d: Fizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d: Buzz\n", i)
		} else if i%7 == 0 {
			fmt.Printf("%d: Bang\n", i)
		} else {
			fmt.Println(i)
		}
	}

}
