//Switch cases avaliam casos de cima para baixo, parando quando um caso for bem-sucedido.

//Este código em Go utiliza o pacote time para imprimir a próxima ocorrência de somente um dia da semana específico (no caso, sábado) em relação à data/hora atual.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")

//A função main começa imprimindo a mensagem "When's Saturday?" na tela usando a função fmt.Println().

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

//Em seguida, usa time.Now().Weekday() para obter o dia da semana atual e armazena o valor retornado na variável today.
//Depois, faz uso da instrução switch para checar se time.Saturday (sábado) ocorre no mesmo dia que today + 0, today + 1 ou today + 2. Dependendo do caso, 
//uma mensagem apropriada é impressa na tela usando fmt.Println().
//Se time.Saturday não ocorrer nos próximos dois dias a partir de today, a mensagem "Too far away." é impressa