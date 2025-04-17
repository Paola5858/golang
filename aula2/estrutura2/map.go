package main

import (
	"fmt"
)

func main() {
	alunoIdade := make(map[string]int)
	alunoIdade["Paola"] = 16
	alunoIdade["Samuel"] = 19
	alunoIdade["Fulano"] = 17
	alunoIdade["Ciclano"] = 18
	fmt.Println("A idade da aluna Paola é:", alunoIdade["Paola"])

	notasAlunos := map[string]float64{
		"Paola": 9.5,
		"Samuel": 8.0,
		"Fulano": 7.5,
		"Ciclano": 9.0,
	}
	for aluno, nota := range notasAlunos {
		fmt.Printf("A nota de %s é: %.2f\n", aluno, nota)
	}
}