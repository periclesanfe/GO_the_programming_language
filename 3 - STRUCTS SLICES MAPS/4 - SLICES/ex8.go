package main

import "fmt"

func main() {
	// É comum acrescentar novos elementos em uma slice, então Go dispõe de uma função padrão para isso o append.

	var s []int
	printSlice(s) // len=0 cap=0 []

	s = append(s, 0) // append funciona tanto para slices nil.
	printSlice(s)    // len=1 cap=1 [0]

	s = append(s, 1) // A slice cresce conforme necessário.
	printSlice(s)    // len=2 cap=2 [0 1]

	s = append(s, 2, 3, 4) // Podemos adicionar mais de um elemento por vez.
	printSlice(s)          // len=5 cap=6 [0 1 2 3 4]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// O primeiro padrâmetro s do append é uma slice do tipo T, e o resto são valores de T para acrescentar ao slice.
// O valor resultante do append é uma slice contendo todos os elementos da slice original mais os valores fornecidos.
// Se a matriz anterior de s for muito pequena para caber os valores, uma matriz gigante será alocada.