//Existem também exportações que é quando levamos o nosso código a buscar por algo de fora, a exemplo o "Pi", que não fora declarado, mas que demonstra resultado, as exportações apresentam sempre as primeiras letras maiúsculas.

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
