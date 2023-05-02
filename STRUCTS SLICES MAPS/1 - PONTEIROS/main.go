package main

import "fmt"

func main() {

	// Go utiliza ponteiros de forma diferente de C. Não há aritmética de ponteiros.

	// Ponteiros guardam na memória um endereço de uma variável.

	i, j := 42, 2701

	// Ponteiros são representados por um asterisco (*) seguido do tipo armazenado pela variável.
	// o tipo *T é um ponteiro para um valor do tipo T. Seu valor zero é nil.
	// O operador & gera um ponteiro para seu operando.
	// O operador * indica valor subjacente do ponteiro

	p := &i         // Ponteiro para i
	fmt.Println(*p) // Lê i através do ponteiro p
	*p = 21         // Define i através do ponteiro p
	fmt.Println(i)  // Imprime o novo valor de i
	p = &j          // Ponteiro para j
	*p = *p / 37    // Divide j através do ponteiro p
	fmt.Println(j)  // Imprime o novo valor de j
}

// Isto é conhecido como "desreferenciamento" ou "indireçionamento".
