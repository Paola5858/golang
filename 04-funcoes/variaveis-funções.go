package main

import "fmt"

// nomeEscola é uma variável de escopo de pacote
// acessível por todas as funções neste pacote
var nomeEscola = "Escola Técnica SENAI - Brilho Divo"

func main() {
	// variáveis locais: existem apenas dentro desta função
	nome := "Paola Fabulosa"
	idade := 16

	mensagem := boasVindas(nome)
	fmt.Println(mensagem)

	status := verificaMaioridade(idade)
	fmt.Println(status)
}
