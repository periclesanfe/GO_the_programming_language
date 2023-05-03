package main

import (
	"fmt"
)

// Chamando aqui a função DIVIDE que pega a/b:
func divide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("ocorreu um erro: %v", r)
		}
	}()
	if b == 0 {
		panic("divisão por zero") //Se ocorrer uma divisão por zero, a função entrará em pânico e a recuperação será feita usando a função "recover", que atualiza o valor do erro para uma mensagem de erro personalizada.
	}
	return a / b, nil
}

// Chamando a função principal:
func main() {
	a := 10
	b := 0
	result, err := divide(a, b)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", result)
	}
}

//Output: Erro: ocorreu um erro: divisão por zero
