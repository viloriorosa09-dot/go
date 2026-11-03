package main

import "fmt"

func main() {

	// calculo de descuento

	var totalCampra uint32
	var descunto float32

	fmt.Print("Por favor diguite el total de su conpra: ")
	fmt.Scan(&totalCampra)

	if totalCampra >= 5000 {
		fmt.Println("felicidades por su compra de ", totalCampra)
		descunto = float32(totalCampra) * 0.20
		fmt.Println("Su descuento por su compra es ", descunto)
		totalCampra = totalCampra - uint32(descunto)
		fmt.Println("a otenido un descuento de el 20% ahora su total a pagra es ", totalCampra)

	} else if totalCampra >= 1000 && totalCampra < 5000 {
		fmt.Println("felicidades por su compra de ", totalCampra)
		descunto = float32(totalCampra) * 0.10
		fmt.Println("Su descuento por su compra es ", descunto)
		totalCampra = totalCampra - uint32(descunto)
		fmt.Println("a otenido un descuento de el 10% ahora su total a pagra es ", totalCampra)
	}

}
