// es como la programacion orientada a odjetos

package main

import "fmt"

type pais struct { // <-- esto es como un clase y sus atributos
	Nombre    string
	Capital   string
	Idioma    string
	Poblacion int
}

func (p pais) detallesDeMiPaies() { // <-- esot es como los metodos
	fmt.Println("el nombre de mi pais es: ", p.Nombre)
	fmt.Println("la capital de mi pais es: ", p.Capital)
	fmt.Println("el idioma mi pais es: ", p.Idioma)
	fmt.Println("la poblacion de mi pais es: ", p.Poblacion)
}

func main() {

	RepublicaDominicana := pais{
		Nombre:  "Republica Dominicana ",
		Capital: "Santo Domingo",
	}

	fmt.Print("Diguite su Idioma aqui: ")
	fmt.Scanln(&RepublicaDominicana.Idioma)

	fmt.Print("Diguite su poblacion aqui: ")
	fmt.Scanln(&RepublicaDominicana.Poblacion)
	fmt.Println("--------------------------------------------------")
	RepublicaDominicana.detallesDeMiPaies()
}
