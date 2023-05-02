// STRUCT LITERAIS

package main

import "fmt"

type Vertex struct {
	X, Y int
}

// A struct literal indica um valor struct recém-alocado, ao enumerar os valores de seus campos,

// você pode listar apenas um subconjunto de campos usando o Name: sintaxe. (E a ordem dos campos nomeados é irrelevante)
// O prefixo especial & constrói um ponteiro para uma struct literal.

var (
	v1 = Vertex{1, 2}  // tem tipo Vertex
	v2 = Vertex{X: 1}  // Y:0 é implícito
	v3 = Vertex{}      // X:0 e Y:0
	p  = &Vertex{1, 2} // tem tipo *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
