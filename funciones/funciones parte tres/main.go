package main

import "fmt"

func sumar(nNumero ...int) int {
	total := 0
	for _, numeros := range nNumero {
		total += numeros
	}
	return total
}

func main() {

	a := sumar(1, 2, 34, 5, 6, 34, 5634, 562, 43, 53246, 432, 56345, 634, 561243, 5134, 5234, 3)

	fmt.Println("la suma de todos los lenguages: ", a)
}
