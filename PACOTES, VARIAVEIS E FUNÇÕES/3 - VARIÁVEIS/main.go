package main

import (
	"fmt"
)
//Para definir uma variavel numérica não necessariamente precisa-se declarar o tipo da variável, pode simplesmente fazer de forma direta:
func main(){
  idade := 50
  fmt.Println(idade) //Output: 50

//Em go, os operadores aritmetricos são deveras simples, onde basta mesmo sem definir função utilizar-se por definição de uma variavel que realiza uma operação! 
  
  x := 5
  y := 3
//SOMA
  soma := x + y
  fmt.Println(soma) // Output: 8
//SUBTRAÇÃO
  subtracao := x - y
  fmt.Println(subtracao) //   Output: 2

//MULTIPLICAÇÃO 
  multiplicacao := x * y
  fmt.Println(multiplicacao) // Output: 15

//DIVISÃO
  divisao := x / y
  fmt.Println(divisao) // Output: 1
  
  modulo := x % y
  fmt.Println(modulo) // Output: 2
}
