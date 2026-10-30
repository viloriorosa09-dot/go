package main

import "fmt"

// funcion valor
func litrosMl(litros float32) float32 {

	return litros / 1000
}

// funcion de referencia en memoria
func mstApies(pies *float32) *float32 {
	*pies = *pies * 3.28084

	return pies
}

func main() {

	var litros float32 = 12.5
	fmt.Println("los litros son ", litros, " a militros ", litrosMl(litros))

	var mts float32 = 5.9

	fmt.Println("los metros  son ", mts)
	fmt.Println("los pies son ", *mstApies(&mts))
}
