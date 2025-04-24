package main
import "fmt"
func main() {
	capitais := map[string]string{
		"SP": "São Paulo",
		"RJ": "Rio de Janeiro",
		"MG": "Minas Gerais",
		"BA": "Bahia",
	}

	for k, v := range capitais {
		fmt.Println("Sigra, Nome", k, v)
	}
}