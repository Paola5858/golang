package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Println("A soma de", a, "e", b, "é:", a+b)
	fmt.Println("A subtração de", a, "e", b, "é:", a-b)
	fmt.Println("A multiplicação de", a, "e", b, "é:", a*b)
	fmt.Println("A divisão de", a, "e", b, "é:", a/b)
	fmt.Println("O resto da divisão de", a, "e", b, "é:", a%b)
	
	logicalOperators()
	}