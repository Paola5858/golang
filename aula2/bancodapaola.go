package main

import "fmt"

func main() {
	var saldo float64
	var valor float64
	var escolha string

	fmt.Println("🎀 Oii, diva! Informe seu saldo inicial:")
	fmt.Scan(&saldo)

	fmt.Println("💰 Agora você pode escolher: 'sacar' ou 'depositar'")
	fmt.Scan(&escolha)

	if escolha == "sacar" {
		fmt.Println("👜 Quanto deseja sacar?")
		fmt.Scan(&valor)

		if valor <= saldo {
			saldo -= valor
			fmt.Println("Saque realizado com sucesso!")
		} else {
			fmt.Println("Ops! Você não tem saldo suficiente, princesa.")
		}
	} else if escolha == "depositar" {
		fmt.Println("💸 Quanto deseja depositar?")
		fmt.Scan(&valor)
		saldo += valor
		fmt.Println("✨ Depósito realizado com sucesso, diva!")
	} else {
		fmt.Println("Eita, esse comando é inválido, tente 'sacar' ou 'depositar'.")
	}

	fmt.Println("💳 Saldo atualizado é:", saldo)
	fmt.Println("💋 Arrasou, linda!")
}
