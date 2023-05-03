//Switch sem uma condição é o mesmo que switch true.

//Essa estrutura pode ser uma maneira limpa para escrever longas cadeias if-then-else.

//Este código em Go utiliza o pacote time para imprimir uma saudação apropriada, baseada no horário do dia em relação à hora atual.

package main

import (
	"fmt"
	"time"
)

func main() {

//A função main começa chamando time.Now() para obter a hora atual e armazena o valor retornado na variável t.

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

//Depois, usa a instrução switch sem uma expressão para comparar o valor de t.Hour() com vários intervalos de horas do dia. 
//Dependendo do caso, uma mensagem apropriada de saudação é impressa na tela usando fmt.Println().
//Se a hora atual for menor que 12, Good morning! é impresso, o que representa a manhã. Caso contrário, se a hora atual for menor que 17, Good afternoon. 
//é impresso, representando a tarde. Se nenhum dos casos acima se aplicar, "Good evening." é impresso, representando a noite.