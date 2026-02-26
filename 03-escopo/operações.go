package main

import "fmt"

// soma retorna a soma de dois inteiros
func soma(x, y int) int {
	return x + y
}

// subtrair retorna a diferença entre dois inteiros
func subtrair(x, y int) int {
	return x - y
}

// multiplicacao retorna o produto de dois inteiros
func multiplicacao(x, y int) int {
	return x * y
}

// divisao retorna o quociente da divisão inteira
// nota: não trata divisão por zero (retorna panic)
func divisao(x, y int) int {
	return x / y
}