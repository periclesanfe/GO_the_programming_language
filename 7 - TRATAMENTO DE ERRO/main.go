// Sintax básica:
package main

import (
	"fmt"
	"os" //Pacote OS, uma biblioteca que permite acessar diversas funções já contidas em GO.
)

// Chamando a FUNC MAIN:
func main() {
	lerArquivo("exemplo.txt")
}

// Define a função que lerá o arquivo, atribuí a ela o valor de stringe diz que sela retornarar ERROR:
func lerArquivo(nomeArquivo string) error {
	arquivo, err := os.Open(nomeArquivo) //os.Open abre arquivos

	if err != nil { // Aqui ocorre o tratamento de EXECEÇÕES utilizando do "nil" ou o "null" de Python, que traz um valor NULO, INEXISTENTE.
		return err //Retorna um ERRO
	}
	defer arquivo.Close() //Aqui o defer garante que o arquivo será lido e fechado mesmo em caso de erro.

	// Lógica para ler o arquivo
	fmt.Println("Arquivo aberto com sucesso!")
	return nil
}

//Output: ./main, demonstrando que houve um erro, para que seja corrijigo.
