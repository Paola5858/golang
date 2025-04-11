package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64

	fmt.Println("💖 Olá, divo(a)! Digite dois números para arrasar nos cálculos: 🎀")
	fmt.Print("🌟 Primeiro número: ")
	fmt.Scan(&num1)
	fmt.Print("🌟 Segundo número: ")
	fmt.Scan(&num2)

	fmt.Println("\n💅✨ Agora veja as operações maravilhosas que podemos fazer com eles: 😘💖")


	fmt.Printf("\n🎀 Soma: %.2f + %.2f = 💖 %.2f\n", num1, num2, num1+num2)
	fmt.Printf("🎀 Subtração: %.2f - %.2f = 💖 %.2f\n", num1, num2, num1-num2)
	fmt.Printf("🎀 Multiplicação: %.2f × %.2f = 💖 %.2f\n", num1, num2, num1*num2)

	if num2 != 0 {
		fmt.Printf("🎀 Divisão: %.2f ÷ %.2f = 💖 %.2f\n", num1, num2, num1/num2)
		fmt.Printf("🎀 Resto da divisão: %.2f %% %.2f = 💖 %.2f\n", num1, num2, float64(int(num1)%int(num2)))
	} else {
		fmt.Println("❌ Divisão por zero? Não, não, meu amor! 💅✨ Vamos evitar essa catástrofe! 😘")
	}

	fmt.Println("\n💋✨ Prontinho, divindade dos cálculos! Você arrasou! 🎀💖")
}
