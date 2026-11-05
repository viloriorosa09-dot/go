//temenos dos formas para hacer mapas un larga y otra corta

package main

import (
	"fmt"
)

func main() {

	// forma lagra

	var diasSemanas = make(map[string]uint8)
	diasSemanas["lunes"] = 1
	diasSemanas["martes"] = 2
	diasSemanas["miercoles"] = 3
	diasSemanas["jueves"] = 4
	diasSemanas["viernes"] = 5
	diasSemanas["sabado"] = 6
	diasSemanas["dimingo"] = 7

	fmt.Println(diasSemanas)

	fmt.Println("--------------------------------------------------------- ")

	//recorer
	for llave, valor := range diasSemanas {
		fmt.Printf("el dia es %v y su numero de semana es %v \n", llave, valor)
	}

	// forma corta

	paisCapital := map[string]string{
		"republica dominicana": "santo domingo",
		"mexico":               "ciunadad de mexico",
		"inglaterra":           "londres",
	}
	fmt.Println(paisCapital)
	fmt.Println("--------------------------------------------------------------------- ")

	for llave, valor := range paisCapital {
		fmt.Printf("el pais es  %v y su capital  es %v \n", llave, valor)
	}

}
