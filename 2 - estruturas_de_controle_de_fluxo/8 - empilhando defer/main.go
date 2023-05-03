//Chamadas de funções adiadas são empurradas por uma pilha. 
//Quando a função retorna, as chamadas de adiamento são executadas na ordem em que a última a entrar é a primeira a sair.

package main

import "fmt"

func main() {
	fmt.Println("counting")

//a mensagem "counting" é impressa na tela	

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

//o loop for é executado 10 vezes, e em cada iteração, o valor atual da variável i é adiado para ser impresso posteriormente usando a palavra-chave defer.

	fmt.Println("done")
}
