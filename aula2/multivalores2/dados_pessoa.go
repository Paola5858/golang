package main

import (
	"fmt"
)

func dadosPessoa(nome string, idade int) (int, string) {
	if idade >= 18 {
		return idade, "maior de idade"
	}
	return idade, "menor de idade"
}

func main() {
	var nome string
	var idade int

	fmt.Print("Digite seu nome diva: ")
	fmt.Scanln(&nome)

	fmt.Print("Digite sua idade maravilhosa: ")
	fmt.Scanln(&idade)

	idadeRecebida, status := dadosPessoa(nome, idade)
	fmt.Printf("Oii %s, você tem %d anos e é %s! 💅\n", nome, idadeRecebida, status)
}
