package main

import "fmt"

func main() {

	// varabel de forma mas larga "palabra reservada var , nombre de la variable, tipo de dato y valor de la variable"
	var nombre string = "Wilber"
	fmt.Println("el nombre es ", nombre)

	/* VARIABLES de forma un pomo mas corto solo se usa si y solo si
	   se ba ainicialzar la variable temenos dos formas una es
	   "la palabra reservada var, nombre de la variable, y el valor de la variable" y la otra es
	   "nombre de la variable dos puntos, signo igual y el valor de la variable" */

	var wilbely = "lixy"
	fmt.Println("el nombre es ", wilbely)
	edad := 1
	fmt.Println("la edad es ", edad)

	// para un bloque de varibles se relaciona entere si

	var (
		nombre2          string = "Wilber"
		edad2            int    = 20
		email            string = "wilber@example.com"
		numeroTelefonico int    = 1234567890
	)

	fmt.Println("el nombre es ", nombre2)
	fmt.Println("la edad es ", edad2)
	fmt.Println("el email es ", email)
	fmt.Println("el numero telefonico es ", numeroTelefonico)

}
