package main

import "fmt"

type Precio struct {
	Modelo      string
	Atiguedad   string
	UsoDelCarro string
}
type Carro struct {
	Marca   string
	Color   string
	PrecioC Precio
}

type pais struct {
	Nombre  string
	Capital struct {
		Nombre    string
		Abitandes uint64
	}
	Idioma    string
	Poblacion uint64
}

func main() {

	var repuclica_dominicana = pais{
		Nombre: "Republica Dominicana",
		Capital: struct {
			Nombre    string
			Abitandes uint64
		}{
			Nombre:    "Santo Domingo",
			Abitandes: 3e6,
		},
		Idioma:    "Español",
		Poblacion: 11e6,
	}

	var toyota = Carro{

		Marca: "Toyota",
		Color: "Negro",
		PrecioC: Precio{
			Modelo:      "Tacoma",
			Atiguedad:   "01/09/2023",
			UsoDelCarro: "Nuevo",
		},
	}

	fmt.Println(repuclica_dominicana)
	fmt.Println(toyota)
}
