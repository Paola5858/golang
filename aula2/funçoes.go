package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "Hello, World!"
	fmt.Println(strings.Contains(greeting, "World")) 
	fmt.Println(strings.ReplaceAll(greeting, "World", "Golang")) 
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "World"))
	fmt.Println(strings.Split(greeting, ", "))
	ages := []int{25, 30, 35}
	sort.Ints(ages) // Sort the ages slice in ascending order
	fmt.Println(ages)
	index := sort.SearchInts(ages, 30) // Search for the index of 30 in the sorted slice
	fmt.Println(index)
	names := []string{"Alice", "Bob", "Charlie"}
	sort.Strings(names) // Sort the names slice in alphabetical order
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "Bob")) // Search for the index of "Bob" in the sorted slice
	

}


