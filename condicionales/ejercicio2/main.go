package main

import "fmt"

func main() {

	var a, b int
	var decimal float64
	var entera int
	var residuo int

	fmt.Print("Diguite le primer valor: ")
	fmt.Scan(&a)
	fmt.Print("Diguite el segundo valor: ")
	fmt.Scan(&b)

	if a < 0 {
		fmt.Println("a es negarivo")
		if b != 0 {
			decimal = float64(a) / float64(b)
			entera = a / b
			residuo = a % b

			if residuo == 0 {
				fmt.Println("Division entera: ", entera)
				fmt.Println("Residuo: ", residuo)
				fmt.Println("Division con decimeles: ", decimal)
				fmt.Println("La Division es exacta")
			} else {
				fmt.Println("Division entera: ", entera)
				fmt.Println("Residuo: ", residuo)
				fmt.Println("Division con decimeles: ", decimal)
				fmt.Println("La Division no es exacta")
			}
		} else {
			fmt.Print("lo siento pero la divicon entre 0 no  esta definida")
		}
	} else if a == 0 {
		fmt.Println("a es cero")
		if b != 0 {
			if residuo == 0 {
				fmt.Println("Division entera: ", entera)
				fmt.Println("Residuo: ", residuo)
				fmt.Println("Division con decimeles: ", decimal)
				fmt.Println("La Division es exacta")
			} else {
				fmt.Println("Division entera: ", entera)
				fmt.Println("Residuo: ", residuo)
				fmt.Println("Division con decimeles: ", decimal)
				fmt.Println("La Division no es exacta")
			}
		} else {
			fmt.Print("lo siento pero la divicon entre 0 no  esta definida")
		}
	} else {
		fmt.Println("a es positivo ")
		if b != 0 {
			if residuo == 0 {
				fmt.Println("Division entera: ", entera)
				fmt.Println("Residuo: ", residuo)
				fmt.Println("Division con decimeles: ", decimal)
				fmt.Println("La Division es exacta")
			} else {
				fmt.Println("Division entera: ", entera)
				fmt.Println("Residuo: ", residuo)
				fmt.Println("Division con decimeles: ", decimal)
				fmt.Println("La Division no es exacta")
			}
		} else {
			fmt.Print("lo siento pero la divicon entre 0 no  esta definida")
		}
	}
}
