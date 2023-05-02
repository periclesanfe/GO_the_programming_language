// Cada programa Go é composto de pacotes.
// Programas começam rodando pelo pacote main.
package main

// Este programa está usando os pacotes com caminhos de importação - que são bibliotecas adicionais ao programa, que trazem funções já criadas, sendo elas nativas ou desenvolvidas por terceiros -   "fmt" e "math/rand".

import (
	"fmt"
	"math/rand"
)

// Por convenção(ordem), o nome do pacote é o mesmo que o último elemento do caminho de importação. Por exemplo, o pacote "math/rand" compreende arquivos que começam com package *rand*(randoom).

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

}
