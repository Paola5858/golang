package main

import "fmt"

func mainVetor() {
 
    ages := [5]int{1, 2, 3, 4, 5}
    names := [5]string{"Paola", "Major", "Fabiano", "Heitor", "Carol"}

    fmt.Println("Ages: ", ages)
    fmt.Println("Names: ", names)

    names[0] = "Paola"
    fmt.Println("Updated Names[0]: ", names[0])

    scores := []int{1, 2, 3, 4, 5}
    fmt.Println("Scores: ", scores)

    scores[0] = 10
    fmt.Println("Updated Scores[0]: ", scores[0])
    fmt.Println(scores, len(scores), cap(scores))

    rangeOne := scores[:2]
    fmt.Println("Range One: ", rangeOne)
    rangeTwo := scores[1:]
    fmt.Println("Range Two: ", rangeTwo)
    rangeThree := scores[:]
    fmt.Println("Range Three: ", rangeThree)

    var superheroes = []string{"Batman", "Superman", "Flash"}
    fmt.Println("Superheroes: ", superheroes)

   
    superheroes = append(superheroes, "Wonder Woman")
    fmt.Println("Updated Superheroes: ", superheroes, len(superheroes), cap(superheroes))

   
    superheroes = append(superheroes[:2], superheroes[3:]...)
    fmt.Println("Updated Superheroes: ", superheroes)
}