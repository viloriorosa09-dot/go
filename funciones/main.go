package main

import "fmt"

const pi = 3.1416

func area(radio float64) float64 {

	return pi * radio * radio
}

func main() {

	var radio float64

	fmt.Print("porfa vor diguite el radio de su circurferencia: ")
	fmt.Scan(&radio)
	fmt.Println("El area de su circulferencia es: ", area(radio))
}
