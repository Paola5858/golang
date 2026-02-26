package main

// boasVindas retorna uma mensagem de boas-vindas personalizada
// usa a variável de pacote nomeEscola
func boasVindas(nome string) string {
	return "Bem-vindo(a), " + nome + "! Você está na " + nomeEscola
}

// verificaMaioridade verifica se a idade é maior ou igual a 18
func verificaMaioridade(idade int) string {
	if idade >= 18 {
		return "Você é maior de idade."
	}
	return "Você é menor de idade."
}
