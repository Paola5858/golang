package main

import "fmt"

// demonstra passagem por valor vs passagem por referência
func modificaPorValor(x int) {
	x = 20
	fmt.Println("dentro da função (valor):", x)
}

func modificaPorReferencia(x *int) {
	*x = 50
	fmt.Println("dentro da função (referência):", *x)
}

func main() {
	// ponteiros: armazenam endereços de memória
	var a int = 10
	var b *int = &a // b aponta para o endereço de a
	
	fmt.Println("valor de a:", a)
	fmt.Println("endereço de a:", b)
	fmt.Println("valor apontado por b:", *b)
	
	// modificando via ponteiro
	*b = 20
	fmt.Println("\napós *b = 20:")
	fmt.Println("valor de a:", a) // a mudou!
	
	// demonstração: passagem por valor
	c := 30
	fmt.Println("\nantes de modificaPorValor:", c)
	modificaPorValor(c)
	fmt.Println("depois de modificaPorValor:", c) // não mudou
	
	// demonstração: passagem por referência
	fmt.Println("\nantes de modificaPorReferencia:", c)
	modificaPorReferencia(&c)
	fmt.Println("depois de modificaPorReferencia:", c) // mudou!
}