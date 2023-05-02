package main

import "fmt"

func main() {

	// Ao 'fatiar', você pode omitir as posições de limite altas ou baixas para usar seus padrões em seu lugar.
// O padrão é igual a zero para o limite baixo, e igual ao comprimento da slice para o limite alto.
// Para a matriz
	s := []int{2, 3, 5, 7, 11, 13}

	// Exemplos de expressões de fatiamento:
	s = s[1:4] 
	fmt.Println(s) // [3 5 7]

	s = s[:2] 
	fmt.Println(s) // [3 5]
 
	s = s[1:]
	fmt.Println(s) // [3, 5, 7, 11, 13]

	// Essas expressões são equivalentes:
	// s = s[0:len(s)]
	// s = s[:len(s)]
	// s = s[0:]
	s = s[:]
	fmt.Println(s) // todas imprimem a slice inteira: [3 5 7 11 13]
}
