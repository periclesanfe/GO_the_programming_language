// A struct é uma coleção de campos

// abaixo um exemplo de struct

package main

import "fmt"

type Vertex struct { // Inicializa uma struct
	X int // X e Y são campos da struct
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2}) // Imprime a struct
}