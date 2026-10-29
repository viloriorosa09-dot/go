package main

import "fmt"

func main() {

	var v1 string = "enero"
	fmt.Println("la variable 1: ", v1, &v1)

	var v2 string = "febrero"
	fmt.Println("la varible 2: ", v2, &v2)

	// apundato al espacion en memoria de v1
	var v3 = &v1

	//mostrando el valor de v1
	fmt.Println("la variable 3: ", *v3, &v3)

}
