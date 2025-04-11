package main

import "fmt"

func main() {
ages := 45
fmt.Println(ages <=50)
fmt.Println(ages >=50)
fmt.Println(ages ==50)
fmt.Println(ages !=50)

if ages <= 30 {
	fmt.Println("Olha só, vc ainda é jovem hein!")
} else if ages < 40 {
	fmt.Println("Olha só, vc é maduro hein!")
} else {
	fmt.Println("Olha só, vc é velho hein!")
}

names := []string{"Paola", "Major", "Fabiano", "Heitor", "Carol"}

for index, name := range names {
	if index == 1 {
		fmt.Println("Essa daqui é a posição:", index, "e esse nosso valor:", name)
		continue
	}
	if index == 3 {
		fmt.Println("sair após essa iteração")
		break
	}
	fmt.Println("Olha só, o nome é: ", name)
}

}