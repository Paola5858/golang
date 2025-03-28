package main

import (
	"fmt"
)

func main() {
    var num1, num2 float64

    fmt.Print("Digite o primeiro número: ")
    fmt.Scan(&num1)
    fmt.Print("Digite o segundo número: ")
    fmt.Scan(&num2)

    fmt.Println("\nResultados das operações aritméticas:")
    fmt.Printf("Soma: %.2f + %.2f = %.2f\n", num1, num2, num1+num2)
    fmt.Printf("Subtração: %.2f - %.2f = %.2f\n", num1, num2, num1-num2)
    fmt.Printf("Multiplicação: %.2f * %.2f = %.2f\n", num1, num2, num1*num2)
    
    if num2 != 0 {
        fmt.Printf("Divisão: %.2f / %.2f = %.2f\n", num1, num2, num1/num2)
        fmt.Printf("Resto da divisão: %.2f %% %.2f = %.2f\n", num1, num2, float64(int(num1)%int(num2)))
    } else {
        fmt.Println("Divisão e resto da divisão não podem ser calculados, pois o divisor é zero!")
    }
}
