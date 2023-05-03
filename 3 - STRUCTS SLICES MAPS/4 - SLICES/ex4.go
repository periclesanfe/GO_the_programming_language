package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	// Uma slice tem tanto um tamanho quanto uma capacidade
	// O comprimento de uma slice é o número de elementos que ela contém.
	// A capacidade de uma slice é o número de elementos na matriz subjacente, contando a partir do primeiro elemento na slice.

	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	// O comprimento e a capacidade de uma slice s podem ser obtidos usando as expressões len(s) e cap(s).
	fmt.Println(len(s)) // 6
	fmt.Println(cap(s)) // 6

	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	s = s[2:]
	printSlice(s) // len=2 cap=4 [5 7]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Você pode estender o comprimento de uma slice "refatiando-a", desde que tenha capacidade suficiente. tente alterar uma das operações da slice no programa de exemplo para estende-la além da sua capacidade e veja o que acontece.
