package main

import "fmt"

type Jogador struct {
	Nome string
	vida int
	nivel int
}
type Carro struct {
	Modelo string
	Ano int
	Velocidade int
	Cor string
}

func exibeCarro(carro Carro) {
	fmt.Println("Modelo:", carro.Modelo)
	fmt.Println("Ano:", carro.Ano)
	fmt.Println("Velocidade:", carro.Velocidade)
	fmt.Println("Cor:", carro.Cor)
}

func exibeDados(jogador Jogador) {
	fmt.Println("Nome:", jogador.Nome)
	fmt.Println("Vida:", jogador.vida)
	fmt.Println("Nível:", jogador.nivel)
}

func main() {
	jogador := Jogador{ Nome: "Paola", vida: 100, nivel: 1}
	exibeDados(jogador)
	carro := Carro{Modelo: "Fusca", Ano: 1970, Velocidade: 80, Cor: "azul"}
	exibeCarro(carro)
}
