package main

import "fmt"

// Uma slice não armazena todos os dados, apenas descreve uma seção de uma matriz subjacente.
// Alterar os elementos de uma slice modifica os elementos correspondentes da sua matriz subjacente.
// Outras slices que compartilham a mesma matriz subjacente verão essas mudanças.

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names) // [John Paul George Ringo]

	a := names[0:2]   // [John Paul]
	b := names[1:3]   // [Paul George]
	fmt.Println(a, b) // [John Paul] [Paul George]

	b[0] = "XXX"
	fmt.Println(a, b)  // [John XXX] [XXX George]
	fmt.Println(names) // [John XXX George Ringo]
}
