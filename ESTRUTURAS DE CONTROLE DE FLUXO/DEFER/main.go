package main

//A estrutura de controle de fluxo em Go é utilizada para executar um bloco de código várias vezes de forma consecutiva,
//O loop "for" é a principal estrutura de repetição em Go. Ele é composto por três componentes: inicialização, condição de continuação e pós-execução.

import "fmt"

func main() {
  // Definindo o tamanho da matriz
  const size = 5

  // Criando a matriz
  var mat [size][size]int

  // Preenchendo a matriz com números aleatórios
  for i := 0; i < size; i++ {
    for j := 0; j < size; j++ {
      mat[i][j] = i + j
    }
  }

  // Imprimindo a matriz gerada
  fmt.Println(mat)

  // Calculando a média dos valores da diagonal principal da matriz
  var diagonalSum float64
  for i := 0; i < size; i++ {
    diagonalSum += float64(mat[i][i])
  }
  diagonalAvg := diagonalSum / float64(size)

  // Imprimindo o resultado do cálculo da média da diagonal principal
  fmt.Printf("A média dos valores da diagonal principal é %.2f\n", diagonalAvg)
}
