package main

import "fmt"

var nomeEscola = "Escola Técnica SENAI - Brilho Divo"

func main() {
    nome := "Paola Fabulosa" // variável local
    idade := 16 // variável local

    mensagem := boasVindas(nome)
    fmt.Println(mensagem)

    status := verificaMaioridade(idade)
    fmt.Println(status)
}
