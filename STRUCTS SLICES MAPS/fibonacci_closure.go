// Exercício: implemente uma função finobacci que retorna uma função (uma closure) que retorna números sucessivos de Fibonacci (0, 1, 1, 2, 3, 5, ...).

package main

import "fmt"

// fibonacci é uma função que retorna uma função que retorna um int.
func fibonacci() func() int {
	// Implemente a função fibonacci.
	return	func() int {
		return 0 // Altere isto para retornar outro valor
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f()) // 0, 1, 1, 2, 3, 5, 8, 13, 21, 34
	}
}
