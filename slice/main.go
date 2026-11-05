package main

import "fmt"

func main() {

	var listaDeAlumnos = make([]string, 10)
	//var enLaLista string

	listaDeAlumnos[0] = "Wilber"
	listaDeAlumnos[1] = "jose"

	for {
		var enLaLista string
		fmt.Print("diguite sun nobre para la lista aqui: ")
		fmt.Scanln(&enLaLista)

		if enLaLista == "" {
			break
		}

		listaDeAlumnos = append(listaDeAlumnos, enLaLista)
	}
	fmt.Println(listaDeAlumnos)
}
