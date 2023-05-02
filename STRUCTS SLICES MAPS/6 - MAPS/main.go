package main

import "fmt"

// Um map mapeia chaves para valores.
// O valor zero de um map é nil. Um map nil não tem chaves, nem chaves podem ser adicionadas.
// A função make retorna um map com um tipo determinado, inicializado e pronto para o uso.

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex) // Inicializando o map
	m["Bell Labs"] = Vertex{40.68433, -74.39967} // Acessando o map e atribuindo valores.
	fmt.Println(m["Bell Labs"]) // Acessando o map e imprimindo valores.
}