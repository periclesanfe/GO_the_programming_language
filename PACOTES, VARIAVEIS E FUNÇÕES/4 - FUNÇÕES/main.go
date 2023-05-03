// A função pode ter zero ou mais argumentos. Neste exemplo, adicione (add) dois parâmetros do tipo int.Observe que o tipo vem após o nome da variável.

package main

import "fmt"

//Pelo fato de ambas serem do tipo "int" pode ser declarada apenas no final:
func soma(x, y int) int {
	return x + y
}
func sub(a, b int) int {
	return a - b
}

//Contudo, declarar pós cada variável não é errado, pelo contrário, dessa forma você pode declarar funcs com variaveis de diferentes tipos:
func div(z int, h int) int {
	return z / h
}
func mult(k int, d int) int {
	return k * d
}

func main() {
	fmt.Println(soma(42, 13)) // Output: 55

	fmt.Println(sub(40, 15)) // Output: 25

	fmt.Println(div(30, 2)) // Output: 15

	fmt.Println(mult(5, 1)) // Output: 5
}

// P.S.: Já existem funções prontas na linguagem que executam tais comandos, como por exemplo a func sum(), para somar;
