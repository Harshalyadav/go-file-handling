package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err :=os.Open("./text.txt")

	b :=make([]byte,4)

	if err != nil{
		fmt.Println(err)
	}

	for{
		 n, error := fileInfo.Read(b)
		if error != nil{
			fmt.Println(err)
			break
		}
		fmt.Println(fileInfo.Name())
		fmt.Println(b[0:n])
		fmt.Println(string(b[0:n]))
	}
	
}