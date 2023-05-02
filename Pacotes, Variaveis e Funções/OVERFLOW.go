package main

import "fmt"

func main() {
    var n, a int
    var op string
    fmt.Print("Digite o valor do seu limite: ")
    fmt.Scanln(&n)
    fmt.Print("Digite o valor da sua variável: ")
    fmt.Scanln(&a)
    fmt.Print("Digite a operação desejada (+, -, *, /): ")
    fmt.Scanln(&op)
    if a < n {
        fmt.Println("OVERFLOW")
    } else {
        switch op {
        case "+":
            fmt.Printf("Resultado: %d\n", n+a)
        case "-":
            fmt.Printf("Resultado: %d\n", n-a)
        case "*":
            fmt.Printf("Resultado: %d\n", n*a)
        case "/":
            if a == 0 {
                fmt.Println("Erro: divisão por zero")
            } else {
                fmt.Printf("Resultado: %d\n", n/a)
            }
        default:
            fmt.Println("Erro: operação inválida")
        }
    }
}
