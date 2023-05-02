package main

import "fmt"

// Constantes são declaradas como variáveis, mas com a palavra-chave const.Constantes podem ser seqüências de caracteres, booleanos, ou valores numéricos.Constantes não podem ser declaradas usando a sintaxe :=.

const Ano = 365

func main() {
	const Hello = "Oi"
	fmt.Println("Hello em português é: ", Hello)
	fmt.Println("Happy", Ano, "Days")

	//Detalhem você não declara o Tipo, ocorre INFERÊNCIA!

	const Truth = true
	fmt.Println("Seguiu as regras?", Truth, "Quantos dias?", Ano, "Então ganha presentes")
}
