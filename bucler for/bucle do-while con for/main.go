// bucle do-while ya que en go no tenemos esto es una forma
//de hacerlo con (for)

package main

import (
	"fmt"
)

func main() {

	var n uint16

	for {
		fmt.Println("1: para suma ")
		fmt.Println("2: para resta ")
		fmt.Println("3: para salir ")
		fmt.Print("Diguite su adcion aqui: ")
		fmt.Scan(&n)

		switch n {
		case 1:
			var n1 float64
			var n2 float64

			fmt.Print("Diguite el primer valor: ")
			fmt.Scan(&n1)
			fmt.Print("Diguite el segundo valor: ")
			fmt.Scan(&n2)
			fmt.Println("la suma de ", n1, " + ", n2, " es: ", n1+n2)
		case 2:
			var n1 float64
			var n2 float64

			fmt.Print("Diguite el primer valor: ")
			fmt.Scan(&n1)
			fmt.Print("Diguite el segundo valor: ")
			fmt.Scan(&n2)
			fmt.Println("la resta de ", n1, " + ", n2, " es: ", n1-n2)

		case 3:
			n = 3
		default:
			fmt.Println("Eres estupido mio o se ase ")
		}

		if n == 3 {
			break
		}
	}
}
