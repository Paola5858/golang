package main

import "fmt"

func pegaNome() (string, string) {
	return "Paola", "Machado"
}

func main() {
	nome, sobrenome := pegaNome()
	fmt.Println("Oii diva, esse é seu nome:", nome)
	fmt.Println("E esse seu sobrenome:", sobrenome)
}

