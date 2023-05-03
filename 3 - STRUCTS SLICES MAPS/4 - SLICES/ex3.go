package main

import "fmt"

func main() {

	// Ao 'fatiar', você pode omitir as posições de limite altas ou baixas para usar seus padrões em seu lugar.
	// O padrão é igual a zero para o limite baixo, e igual ao comprimento da slice para o limite alto.
	// Para a matriz
	s := []int{2, 3, 5, 7, 11, 13}

	// Exemplos de expressões de fatiamento:
	t := s[1:4]
	fmt.Println(t) // [3 5 7]

	u := s[:2]
	fmt.Println(u) // [2 3]

	v := s[1:]
	fmt.Println(v) // [3, 5, 7, 11, 13]

	// Essas expressões são equivalentes:
	// s = s[0:len(s)]
	// s = s[:len(s)]
	// s = s[0:]
	x := s[:]
	fmt.Println(x) // todas imprimem a slice inteira: [3 5 7 11 13]
}
