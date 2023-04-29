package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err :=os.ReadFile("./text.txt")

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(fileInfo)
	fmt.Println(string(fileInfo))
	
}