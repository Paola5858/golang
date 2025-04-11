package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
    // Manipulação de strings
    greeting := "Olá, Mundo!"
    fmt.Println("A saudação contém 'Mundo'? ", strings.Contains(greeting, "Mundo"))
    fmt.Println("Substituindo 'Mundo' por 'Golang':", strings.ReplaceAll(greeting, "Mundo", "Golang"))
    fmt.Println("Saudação em maiúsculas:", strings.ToUpper(greeting))
    fmt.Println("Índice da palavra 'Mundo':", strings.Index(greeting, "Mundo"))
    fmt.Println("Dividindo a saudação:", strings.Split(greeting, ", "))

    // Trabalhando com slices de inteiros
    ages := []int{25, 30, 35}
    sort.Ints(ages)
    fmt.Println("Idades ordenadas:", ages)
    index := sort.SearchInts(ages, 30)
    fmt.Println("A idade 30 está no índice:", index)

    // Trabalhando com slices de strings
    names := []string{"Paola", "Major", "Fabiano"}
    sort.Strings(names)
    fmt.Println("Nomes ordenados:", names)
    fmt.Println("O índice do nome 'Paola' é:", sort.SearchStrings(names, "Paola"))

    // Laços de repetição
    fmt.Println("\nContagem regressiva:")
    x := 5
    for x > 0 {
        fmt.Printf("Faltam %d...\n", x)
        x--
        }
    fmt.Println("Lançar!")

    fmt.Println("\nContando até 5:")
    for i := 0; i < 5; i++ {
        fmt.Printf("Número: %d\n", i)
    }

    fmt.Println("\nExibindo nomes com índice:")
    for i := 0; i < len(names); i++ {
        fmt.Printf("Nome %d: %s\n", i, names[i])
    }

    fmt.Println("\nExibindo nomes com índice e valor:")
    for index, value := range names {
        fmt.Printf("O índice é %d e o nome é %s\n", index, value)
    }

	superherois := []string{"Batman", "Superman", "Mulher Maravilha"}
	fmt.Println("\nSuper-heróis:")

	for index, hero := range superherois {
		fmt.Printf("O número do Super-herói é %d e o nome é %s\n", index, hero)
	}

	for i:= 0; i < len(superherois); i++ {
		fmt.Printf("O número do Super-herói é %d e o nome é %s\n", i, superherois[i])
	}
}