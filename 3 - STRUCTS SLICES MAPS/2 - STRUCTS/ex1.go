// campos struct são acessados através de um ponto

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2} // Define v como uma struct com os valores 1 e 2
	v.X = 4 // Acessa o campo X da struct v e atribui o valor 4
	fmt.Println(v.X) // Imprime o valor do campo X da struct v
}
