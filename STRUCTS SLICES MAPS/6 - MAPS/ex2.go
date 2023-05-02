package main

import "fmt"

func main() {

	// Inserir ou atualizar um elemento no map m:
	// m[key] = elem
	
	m := make(map[string]int)
	
	// Recuperar um elemento:
	// elem = m[key]
	
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"]) 

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// Deletar um elemento:
	// delete(m, key)

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// Testar se uma chave está presente com uma atribuição de dois valores:
	// elem, ok = m[key]
	// Se key estiver em m, ok será true. Se não, ok será false.
	// Se key não está no map elem tem valor zero para elemento do tipo map.

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
	
}
