package main

import (
	"fmt"
	"time"
)

func main() {

	// PROCESSO
	// 1. Uma instância de um programa em execução é chamada de processo.
	// 2. Um processo fornece ambiente para um programa executar.

	// THREAD
	// 1. São as menores unidades de execução que a CPU aceita.
	// 2. Um processo pode ter uma ou mais threads.

	// GOROUTINE
	// 1. Uma go routine é um encadeamento leve (logicamente um encadeamento de execução) gerenciado pela runtime do GO.
	// 2. Uma go routine é uma função que é executada em paralelo com outras funções.
	// 3. Para iniciar um novo goroutine executando, adicione a palavra-chave go andes da chamada da função.
	go printNumbers()
	go printLetters()

	time.Sleep(time.Second)

	fmt.Println("Fim do programa")
}

func printNumbers() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for i := 'A'; i < 'K'; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

// Go routine são gerenciadas pelo GOScheduler que é parte da runtime do GO.
// Go routine faz o papel de uma thread, mas não é uma thread. já que existe apenas no espaço virtual do GoLang e não no sistema.
// Também possui os estados de Wainting/Blocked, Executable/Runnable e Running.
