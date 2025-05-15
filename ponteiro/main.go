package main 

import "fmt"

func mudar(x int) {
	x = 20
	fmt.Println("Valor de x dentro da função mudar:", x)
	z := 30
	fmt.Println("Valor de z dentro da função mudar:", z)
}

func main() {
	var a int = 10
	var b *int = &a
	fmt.Println("Valor de b:", b)
	*b = 20
	fmt.Println("Valor de a:", a)

	c := 30
	fmt.Println("Valor de c:", c)
}