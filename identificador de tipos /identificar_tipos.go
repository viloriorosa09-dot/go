package main

import (
	"fmt"
	"reflect"
)

func main() {

	nombre := "Wilber"
	edad := 20
	peso := 182.0
	altura := 1.83

	fmt.Println("el tipo de dato es ", reflect.TypeOf(nombre))
	fmt.Println("el tipo de dato es ", reflect.TypeOf(edad))
	fmt.Println("el tipo de dato es ", reflect.TypeOf(peso))
	fmt.Println("el tipo de dato es ", reflect.TypeOf(altura))

	// otra forma de identificar el tipo de dato
	fmt.Printf("el tipo de dato es %T y el nombre es %v \n", nombre, nombre)
	fmt.Printf("el tipo de dato es %T y la edad es %v \n", edad, edad)
	fmt.Printf("el tipo de dato es %T y el peso es %v \n", peso, peso)
	fmt.Printf("el tipo de dato es %T y la altura es %v \n", altura, altura)

}
