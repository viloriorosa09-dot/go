package main

import "fmt"

func saludar(saldo string) string {
	return "hola como esta señor" + saldo
}

func main() {

	nombre := " Wilber"
	apellido := " Vilorio rosa"

	fmt.Println(saludar(nombre), apellido)
}
