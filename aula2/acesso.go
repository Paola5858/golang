package main

import (
	"fmt"
)

func acesso() {
	var usuario, senha string

	fmt.Println("💖 Bem-vindo ao sistema mais glamouroso do universo! 🌟")
	fmt.Print("🎀 Digite seu usuário: ")
	fmt.Scan(&usuario)

	fmt.Print("🎀 Agora, digite sua senha: ")
	fmt.Scan(&senha)


	if usuario == "admin" && senha == "1234" {
		fmt.Println("\n💋✨ Acesso PERMITIDO, divo(a)! Você brilhou! 💖🎀✨")
	} else {
		fmt.Println("\n❌ Acesso NEGADO, meu bem! Tente novamente! 💅")
	}

	fmt.Println("💎💖 Obrigado por usar nosso sistema fabuloso! 😘✨")
}
