package main

import "fmt"

func main() {

	var diaSemana uint8

	fmt.Print("diguite el dia de la semana que es: ")
	fmt.Scan(&diaSemana)

	switch diaSemana {

	case 1:
		fmt.Println("El dia de la semana es Lunes.")
	case 2:
		fmt.Println("El dia de la semana es Martes.")
	case 3:
		fmt.Println("El dia de la semana es Miercoles.")
	case 4:
		fmt.Println("El dia de la semana es Jueves.")
	case 5:
		fmt.Println("El dia de la semana es Viernes.")
	case 6:
		fmt.Println("El dia de la semana es Sabado.")
	case 7:
		fmt.Println("El dia de la semana es Domingo.")
	default:
		fmt.Println("Usted vive en marte ese dia de la semana no existe jodase")
	}

}
