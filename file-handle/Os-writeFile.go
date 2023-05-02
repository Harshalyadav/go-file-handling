package main

import (
	"fmt"
	"os"
)

func main() {

	fileInfo, err := os.OpenFile("./1text.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		fmt.Println(err)
	}

	defer fileInfo.Close()
	_, err = fileInfo.WriteString("Hop you had a good day!")

}
