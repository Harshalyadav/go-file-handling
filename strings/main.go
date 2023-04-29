package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("......Learn String........")

	learn1 := "Hi I'm learning advance go language"
	learn2 := "Learning go is easy and fun"
	learn3 := "Learning"
	fmt.Println("......Container......")
	result := strings.Contains(learn1,learn2)
	fmt.Println(".....Count...........")
	result1 := strings.Count(learn1,learn3)
	result2 := strings.EqualFold(learn1,learn2)
	fmt.Println(result)
	fmt.Println(result1)
	fmt.Println(result2)
}