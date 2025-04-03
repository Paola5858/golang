package main

import "fmt"

func main() {
    var ages = [5]int{1, 2, 3, 4, 5}
    nomes := [5]string{"Paola", "Major", "Fabiano", "Heitor", "Carol"}

    fmt.Println(ages)
    fmt.Println(nomes)

    nomes[0] = "Paola"
    fmt.Println(nomes[0])
}
