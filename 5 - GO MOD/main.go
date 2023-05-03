// GO MOD
// Permite o controle de dependências de um projeto, também como versões.
// Independe do caminho GOPATH e não precisa de um diretório src.
// O arquivo go.mod é criado automaticamente quando se usa o comando go mod init.
// Para teste: go mod init github.com/username/projectname
// Para adicionar uma dependência: go get github.com/username/projectname

// go mod init github.com/codeedu/myuuid
// go get github.com/google/uuid@latest

package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// generate uuid v4
	uuid := uuid.New()
	fmt.Println(uuid)
}

// go mod tidy - remove dependências não utilizadas e instala as que estão faltando
// go mod graph - mostra o grafo de dependências
// go mod vendor - cria um diretório vendor com as dependências utilizadas na máquina local.
// go install nomeDoPacore@versao - compila e instala o pacote no diretório pkg
