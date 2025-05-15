package main 

import "fmt"

func main() {
	var a int = 10
	var b *int = &a
	fmt.Println("Valor de b:", b)
	*b = 20
}