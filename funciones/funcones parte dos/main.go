package main

import "fmt"

func calculadora(a int, b int) (suma int, resta int, mult int, div float64) {
	suma = a + b
	resta = a - b
	mult = a * b

	if b == 0 {
		fmt.Errorf("La divicion entre cero no esta definida :( )")

	} else {
		div = float64(a) / float64(b)
	}

	return
}

func main() {

	var a int
	var b int

	// leer datos por teclado

	fmt.Print("Diguite el primer valor: ")
	fmt.Scanln(&a)
	fmt.Print("Diguite el segundo valor: ")
	fmt.Scanln(&b)

	suma, resta, mult, div := calculadora(a, b)

	fmt.Println("\nla suma es: ", suma)
	fmt.Println("la resta es: ", resta)
	fmt.Println("La multiplicacion es: ", mult)
	fmt.Println("la divicion es: ", div)
}
