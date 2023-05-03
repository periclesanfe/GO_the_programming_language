// SLICES LITERAIS

package main

import "fmt"

func main() {
	
	// Uma slice literal é um array literal sem o comprimento.
	// Esta é uma slice literal:
	// [3]bool{true, true, false}
	// E isso cria a mesma matriz como acima, e depois constrói uma slice que aponta para ele:
	q := []int{2, 3, 5, 7, 11, 13}
	
	fmt.Println(q) // [2 3 5 7 11 13]
	
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r) // [true false true true false true]

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s) // [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
}
