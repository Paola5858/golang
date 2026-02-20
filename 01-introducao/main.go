package main

import "fmt"

func main() {

	fmt.Println("Oii, Golang! Preparada pra brilhar? 💋✨")
	var nomeOne string = "Paola"
	var nomeTwo = "Major"
	var nomeTree string
	fmt.Println("Elenco original:", nomeOne, nomeTwo, nomeTree)
	nomeTree = "Fabiano"
	nomeTwo = "Heitor"
	fmt.Println(" - Após a reviravolta:", nomeTree, nomeTwo)
	nomeFour := "Carol"
	fmt.Println(" - Novo personagem entra na história:", nomeFour)
	var scoreOne float32 = 0.0
	scoreTwo := 10.0
	fmt.Println(" -  Placar atualizado:", scoreOne, scoreTwo)
	const pi float32 = 3.14159
	const nome string = "Paola"
	fmt.Println(" -  Constantes matemáticas e de identidade:", pi, nome)
	fmt.Println(" -  Fim da execução! Agora é sua vez de brilhar no Go! 🌟")
}
