package main

import (
    "fmt"
)

func main() {
    fmt.Println("Digite uma senha: ")
    var password string
    var attempts int
    var locked_out bool

    do {
        fmt.Scanln(&password)
        if password == "senha-123" {
            fmt.Println("Bem-vindo!")
            break
        } else {
            attempts++
            if attempts >= 3 {
                fmt.Println("Você excedeu o número máximo de tentativas.")
                locked_out = true
                break
            }
            fmt.Println("Senha incorreta. Tente novamente: ")
        }
    } while (!locked_out)

    if !locked_out {
        // Código a ser executado quando a senha estiver correta.
        fmt.Println("Acesso concedido.")
    }
}

//O do-while é usado para solicitar a senha e verificar se está correta ou se o usuário excedeu o número máximo de tentativas. 
//O loop é interrompido com break quando o usuário faz login com sucesso ou quando o número máximo de tentativas é alcançado.