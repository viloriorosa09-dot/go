package main

import (
	"fmt"
)

func factorial(n uint32) uint32 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

}

func main() {
	var factorial1 uint32

	fmt.Print("Pro favor diguite le numero: ")
	fmt.Scan(&factorial1)
	fmt.Println("su factorial es ", factorial(factorial1))
}
