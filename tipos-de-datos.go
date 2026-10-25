package main

import "fmt"

func saludo(saludo string) string {
	return "hola como estas " + saludo
}

func sumar(a int, b int) int {
	return a + b
}

func main() {

	edad := 20
	nombre := " Wilber"
	var peso float32 = 182.0
	var esEstudiante bool = true
	altura := 1.83

	fmt.Println(sumar(22, 42))
	fmt.Println(saludo(nombre))
	fmt.Println("peso: ", peso)
	fmt.Println("altura: ", altura)
	fmt.Println("es estudiante: ", esEstudiante)
	fmt.Println("mi edad es: ", edad)

}
