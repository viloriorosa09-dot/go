package main

import (
	"fmt"
	"strconv" //<-- para conversión de tipos.
)

func main() {
	// convirtiendo un texto a buleano.

	var textoBool1 = "treu"
	var boleanoReal, error = strconv.ParseBool(textoBool1)

	fmt.Printf("esto era texto %v y ahora es un boleano %v y lo confirmo %T \n", textoBool1, boleanoReal, error)

	// para hacer la oprecion inversa, es decir convertir un boleano a texto, se hace de la siguiente manera.

	boleano := true
	var textoBool2 = strconv.FormatBool(boleano)

	fmt.Printf("esto era un boleano %v y ahora es un texto %v y lo confirmo %T \n", boleano, textoBool2, textoBool2)

	// convertir un numaro a texto con strconv.

	numero := 33
	var textoNumero = strconv.Itoa(numero)

	fmt.Printf("esto era un numero %v y ahora es un texto %v y lo confirmo %T \n", numero, textoNumero, textoNumero)

	// para convertir un texto a numero con strconv.

	textoNumero2 := "22"
	var numero2, error2 = strconv.Atoi(textoNumero2)

	fmt.Printf("esto era un texto %v y ahora es un numero %v y lo confirmo %T \n", textoNumero2, numero2, error2)

	// para convertir un numero flotante a entero.

	decimal := 9.99
	numeroEntero := int(decimal)

	fmt.Printf("esto era un numero decimal %v y ahora es un numero entero %v y lo confirmo %T \n", decimal, numeroEntero, numeroEntero)
}
