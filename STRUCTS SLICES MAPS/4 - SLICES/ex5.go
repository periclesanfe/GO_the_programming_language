package main

import "fmt"

func main() {
	var s []int

	// Uma slice nula tem o comprimento e a capacidade igual a 0.
	
	fmt.Println(s, len(s), cap(s)) // [] 0 0

	// O valor zero de uma slice é nil.

	if s == nil {
		fmt.Println("nil!")
	}
}