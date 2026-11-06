package main

import "fmt"

type Pais struct {
	Nombre    string
	Capital   string
	Idioma    string
	Poblacion uint64
}

func main() {

	var paises = make([]Pais, 0)

	continua := "N"

	for {

		fmt.Print("Quiere ingresar un pais (Y/N): ")
		fmt.Scanln(&continua)

		if continua != "Y" {
			break
		}

		mipais := Pais{}

		fmt.Print("Diguite el pais: ")
		fmt.Scanln(&mipais.Nombre)
		fmt.Print("Diguite la capital: ")
		fmt.Scanln(&mipais.Capital)
		fmt.Print("Diguite el idioma: ")
		fmt.Scanln(&mipais.Idioma)
		fmt.Print("Diguite la poblacion: ")
		fmt.Scanln(&mipais.Poblacion)
		paises = append(paises, mipais)
	}

	for _, pais := range paises {
		fmt.Println("-------------------------------------------------------")
		fmt.Println(pais)
	}
}
