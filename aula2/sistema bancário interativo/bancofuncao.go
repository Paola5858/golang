package main

import (
	"fmt"
)

func saudacao(nome string) {
	fmt.Printf("\n✨ Oii, %s! Bem-vinda ao nosso sistema bancário 💋\n\n", nome)
}

func somar(a int, b int) int {
	return a + b
}

func verSaldo() {
	fmt.Println("Seu saldo é de R$ 300,00")
}
func sacar() {
	fmt.Println("Você sacou R$ 20,00")
}
func depositar() {
	fmt.Println("Você depositou R$ 20,00")
}
func transferir() {
	fmt.Println("Você transferiu R$ 20,00")
}

func mostrarMenu() {
	fmt.Println("🎀 Escolha uma opção:")
	fmt.Println("1 - Ver saldo")
	fmt.Println("2 - Sacar")
	fmt.Println("3 - Depositar")
	fmt.Println("4 - Transferir")
	fmt.Println("5 - Sair")
	fmt.Print("💖 Digite sua opção: ")
}

func main() {
	var nome string
	fmt.Print("Informe seu nome para começar: ")
	fmt.Scan(&nome)

	saudacao(nome)
	mostrarMenu()

	var opcao int
	fmt.Scan(&opcao)

	switch opcao {
	case 1:
		verSaldo()
	case 2:
		sacar()
	case 3:
		depositar()
	case 4:
		transferir()
	case 5:
		fmt.Println("Saindo do sistema... volte sempre!")
	default:
		fmt.Println("Opção inválida. Tente novamente.")
	}

	resultado := somar(7, 3)
	fmt.Printf("\nTeste de função de soma: 7 + 3 = %d 💕\n", resultado)

	fmt.Println("\nObrigada por usar nosso sistema, diva!")
	fmt.Println("xoxo 💋💋💋")
}
