package main
import (
	"fmt"
)

func dividir (dividendo int , divisor int) (int, string) {
	if divisor == 0 {
		return 0, "Divisão por zero não é permitida."
	}
	resultado := dividendo / divisor
	return resultado, ""
}

func operaçãoBasica(a int, b int) (int, int, int) {
	soma := a + b
	subtracao := a - b
	multiplicacao := a * b
	return soma, subtracao, multiplicacao
}
func main() {
	resultado, erro := dividir(10, 2)
	if erro != "" {
		fmt.Println("Erro:", erro)
	} else {
		fmt.Println("Oii diva, seu resultado deu:", resultado)
	}
	soma, subtracao, multiplicacao := operaçãoBasica(10, 5)
	fmt.Println("Essa é sua soma diva:", soma)
	fmt.Println("Essa é sua subtração diva:", subtracao)
	fmt.Println("Essa é sua multiplicação diva:", multiplicacao)
}
