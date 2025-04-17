package main

import "fmt"

func main() {
    var numero int
    fmt.Println("Olá! Digite um número:")
    if _, err := fmt.Scan(&numero); err != nil {
        fmt.Println("Erro ao ler o número:", err)
        return
    }
    fmt.Printf("O número digitado foi: %d\n", numero)
}