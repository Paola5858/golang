package main

import "fmt"

// Jogador representa um jogador no sistema
type Jogador struct {
	Nome  string
	Vida  int
	Nivel int
}

// Carro representa um veículo com suas propriedades
type Carro struct {
	Modelo     string
	Ano        int
	Velocidade int
	Cor        string
}

// ExibirStatus é um método do tipo Jogador
func (j Jogador) ExibirStatus() {
	fmt.Printf("\nJogador: %s\n", j.Nome)
	fmt.Printf("Vida: %d\n", j.Vida)
	fmt.Printf("Nível: %d\n", j.Nivel)
}

// Acelerar é um método que modifica o estado do carro
func (c *Carro) Acelerar(incremento int) {
	c.Velocidade += incremento
	fmt.Printf("%s acelerou! Velocidade atual: %d km/h\n", c.Modelo, c.Velocidade)
}

// ExibirDetalhes mostra informações do carro
func (c Carro) ExibirDetalhes() {
	fmt.Printf("\nCarro: %s (%d)\n", c.Modelo, c.Ano)
	fmt.Printf("Cor: %s\n", c.Cor)
	fmt.Printf("Velocidade: %d km/h\n", c.Velocidade)
}

func main() {
	// criando structs com valores iniciais
	jogador := Jogador{
		Nome:  "Paola",
		Vida:  100,
		Nivel: 1,
	}
	
	carro := Carro{
		Modelo:     "Ferrari F8",
		Ano:        2023,
		Velocidade: 0,
		Cor:        "Vermelho",
	}
	
	// usando métodos
	jogador.ExibirStatus()
	carro.ExibirDetalhes()
	
	fmt.Println("\n--- Ações ---")
	carro.Acelerar(50)
	carro.Acelerar(100)
}