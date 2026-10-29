package main

import "fmt"

func esAdulto(edad int) bool {
	return edad >= 18
}

func main() {

	verdadero := true
	var falso bool

	fmt.Print(verdadero, falso)

	if esAdulto(20) {
		fmt.Println("Usted puede entrar al var.")

	} else {
		fmt.Println("Usted no puede entrar al var.")
	}

}
