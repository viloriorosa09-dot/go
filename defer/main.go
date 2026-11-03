package main

import "fmt"

func desconerctar() {
	fmt.Println("Sea desconectado de la base de datos")
}

func leer() {
	fmt.Println("sean leidao los reguistros de la base de datos ")
}

func actualizar() {
	fmt.Println("sean actualizados los reguitros de la base de datos ")
}
func conectar() {
	fmt.Println("sean conectado la base de datos ")
}
func main() {

	conectar()
	defer desconerctar()
	leer()
	actualizar()

}
