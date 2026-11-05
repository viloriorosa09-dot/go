package main

import (
	"fmt"
)

func sumaExpecifica(a int8, b int8) uint16 {
	return (uint16(a) + uint16(b))
}

func main() {

	//variables
	resultado := sumaExpecifica(55, 43)
	var altula float32 = 2.11
	var queso float64 = 3.2222232222223232233232232323

	//salida
	fmt.Println(resultado, altula, queso)

}
