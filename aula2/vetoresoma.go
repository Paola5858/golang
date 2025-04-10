package main

import "fmt"

func main() {
    var numeros [5]int
    soma := 0

    fmt.Println("Digite 5 números inteiros 🎯:")
    for i := 0; i < len(numeros); i++ {
        fmt.Printf("Número %d: ", i+1)
        fmt.Scan(&numeros[i])
        soma += numeros[i]
    }

    fmt.Printf("\nA soma dos números digitados é: %d 💖\n", soma)
    fmt.Println("Arrasou, diva! Código finalizado com sucesso 📝✨")
}
