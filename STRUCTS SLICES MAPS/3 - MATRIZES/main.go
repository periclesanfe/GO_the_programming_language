// Matrizes em GO funciona de forma semelhante a C, mas de forma diferente de C++, Java e C#.
// Uma matriz é uma sequência numerada de elementos de um único tipo.
// O número de elementos é chamado de comprimento e é nunca negativo.
// O comprimento de uma matriz m pode ser descoberto usando a função len(m).

package main

import "fmt"

// O tipo [n]T é uma matriz de n valores de tipo T.

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) // Hello World
	fmt.Println(a)          // [Hello World]

	// Uma matriz literal é a lista de valores entre chaves.
	// ... permite que o compilador conte o número de elementos.
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[0], primes[1], primes[2], primes[3], primes[4], primes[5]) // 2 3 5 7 11 13
	fmt.Println(primes) // [2 3 5 7 11 13]
}
