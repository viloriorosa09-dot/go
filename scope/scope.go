package main

import "fmt"

var nivel1 = "esta es mi varible del nuvel 1"

func main() {
	fmt.Println(nivel1)

	var nivel2 = "esta es mi varible del nivel 2"
	fmt.Println(nivel2)
	{
		var nivel3 = "esta es mi variable del nivel 3"
		fmt.Println(nivel3)
	}

}
