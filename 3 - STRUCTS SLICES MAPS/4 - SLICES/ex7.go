// SLICES DE SLICES

package main

import (
	"fmt"
	"strings"
)

func main() {
	// As slices podem conter qualquer tipo, incluindo outras slices.

	// Criando um tabuleiro de jogo da velha.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// rodadas dos players
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	// imprimindo o tabuleiro
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
