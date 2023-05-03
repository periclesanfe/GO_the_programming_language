package main

import "fmt"

// fibonacci é uma função que retorna uma função que retorna um int.
func fibonacci() func() int {
	a, b := 0, 1 // inicializa as variáveis para a série de Fibonacci
	return func() int {
		result := a
		a, b = b, a+b // atualiza os valores para a próxima iteração
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f()) // imprime os 10 primeiros números da série de Fibonacci
	}
}
