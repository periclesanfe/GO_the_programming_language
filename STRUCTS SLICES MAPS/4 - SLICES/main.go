// Uma matriz tem tamanho fixo. Uma slice, por outro lado, é dinamicamente redimensionada, uma visão flexível dos elementos de uma matriz. Na prática, as slices são muito mais comuns do que as matrizes.
// O tipo []T é uma slice com elementos do tipo T.
// Uma slice é formada por uma especificação de dois indices, um limite menor e maior, separados por dois pontos:

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// A sintaxe a[low: high] seleciona um intervalo meio-aberto que inclui o primeiro elemento, mas exclui o último.
	// A expressão a seguir cria uma slice que inclui os elementos de 1 à 3 de a:

	var s []int = primes[1:4]
	fmt.Println(s) // [3 5 7]
}
