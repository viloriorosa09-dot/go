// bucle while en go reprecentado con for ya que no tenemos esos bucles

package main

import "fmt"

func main() {

	var n int
	//variable iteradora --> i
	i := 1

	fmt.Print("Diguite el numero de la tabla que desa aprender: ")
	fmt.Scan(&n)

	for i <= 12 {
		fmt.Println(n, " x ", i, " = ", n*i)
		i++
	}
}
