package main

import "fmt"

func main() {

	var juego = []string{"wilber ", "valia", "pedro ", "yota", "juana"}

	juego = append(juego[:2], juego[3:]...)

	fmt.Print(juego)
}
