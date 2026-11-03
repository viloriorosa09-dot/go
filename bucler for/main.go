// bucler for natural en cualquier lenguage

package main

import "fmt"

func main() {

	var n int

	fmt.Print("por favor diguite el numero de la tabla que decea: ")
	fmt.Scan(&n)

	for i := 1; i <= 12; i++ {
		fmt.Println(n, " x ", i, " = ", n*i)

	}
}
