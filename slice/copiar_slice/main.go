// temos dos for mas para hacerlo

package main

import "fmt"

func main() {

	// primera forma
	var figuras = make([]string, 3)
	figuras1 := []string{"", "", ""}

	figuras[0] = "cuadrado"
	figuras[1] = "circulo"
	figuras[2] = "triangulo"

	copy(figuras1, figuras)

	fmt.Println(figuras)
	fmt.Println(figuras)
	// segunda forma

	var nombres = []string{"wilber", "juna", "rosa"}
	nombres2 := append([]string{}, nombres...)
	fmt.Println(nombres)
	fmt.Println(nombres2)

}
