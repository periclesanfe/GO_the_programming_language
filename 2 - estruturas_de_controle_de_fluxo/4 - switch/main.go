//Uma instrução switch é uma forma mais curta de escrever uma sequência de declarações if - else. 
//Ele executa o primeiro caso cujo valor é igual à expressão de condição.

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}