package main

import "fmt"

type caro struct {
	Marca  string
	Color  string
	Modelo string
	Precio uint64
}

func main() {

	var toyota = caro{
		Marca:  "Toyota",
		Color:  "Blanco",
		Modelo: "Tacoma",
		Precio: 3e6,
	}

	var honda = new(caro)
	honda.Marca = "Honda"
	honda.Color = "Negro"
	honda.Modelo = "CRV"
	honda.Precio = 2e6

	fmt.Println("La marca es: ", toyota.Marca)
	fmt.Println("El color es: ", toyota.Color)
	fmt.Println("El modelo es: ", toyota.Modelo)
	fmt.Println("El precio es: ", toyota.Precio)
	fmt.Println("---------------------------------------")
	fmt.Println("La marca es : ", honda.Marca)
	fmt.Println("El color es : ", honda.Color)
	fmt.Println("El modelo es : ", honda.Modelo)
	fmt.Println("El precio es : ", honda.Precio)

}
