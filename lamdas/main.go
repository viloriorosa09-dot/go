package main

import "fmt"

const PI float64 = 3.1416

func circulo(radio float64) (area func() float64, perimetro func() float64) {

	area = func() float64 {
		return PI * radio * radio
	}

	perimetro = func() float64 {
		return 2 * PI * radio
	}
	return
}

func main() {
	var radio uint32

	fmt.Print("Por favor diguite el radio: ")
	fmt.Scan(&radio)

	area, perimetro := circulo(float64(radio))

	fmt.Println("el area es:", area())
	fmt.Println("el perimetro es: ", perimetro())

}
