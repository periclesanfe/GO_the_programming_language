//A declaração if parece com os laços for; a expressão não precisa ser cercada de ( ) mas os chaves { } são obrigatórios.

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {

//A função sqrt verifica se o argumento fornecido é um número negativo e, se for, 
//calcula a raiz quadrada do número positivo correspondente e acrescenta a letra "i" para indicar que o resultado é um número imaginário.

	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))

// Se o número for não negativo, a função math.Sqrt do pacote math é usada para calcular a raiz quadrada numérica e a resposta é convertida em uma string.

}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}