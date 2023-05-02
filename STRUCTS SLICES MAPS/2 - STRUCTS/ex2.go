// Campos de estrutura podem ser acessados através de um ponteiro de estrutura

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// Para acessar o campo X de uma struct, quando tivernos o ponteiro p da strut, podemos escrever (*p).X

func main() {
	v := Vertex{1, 2}
	p := &v
	// (*p).X = 1e9
	p.X = 1e9 // Go permite escrever apenas p.X ao invés de (*p).X, sem referência explicia ao ponteiro.
	fmt.Println(v)
}
