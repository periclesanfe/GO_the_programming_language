// MAPS LITERAIS

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// Você pode declarar um map literal assim:
var m = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

// Se o tipo de nível superior é apenas um nome do tipo, você pode omiti-lo a partir dos elementos do lietral.
// Exemplo:
// var m = map[string]Vertex{
// 	"Bell Labs": {40.68433, -74.39967},
// 	"Google":    {37.42202, -122.08408},
// }

func main() {
	fmt.Println(m) // Imprimindo o map literal.
}
