//Variáveis ​​declaradas dentro de uma instrução if curta também estão disponíveis dentro dos blocos else.
//(Ambas as chamadas pow são executadas e retornam antes da chamada fmt.Println no começo do main.)

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {

//A função pow recebe três argumentos do tipo float64 e retorna um valor float64. 

	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}	
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

//Na função principal (main), há uma chamada para a função fmt.Println() que imprime o resultado de duas chamadas da função pow().