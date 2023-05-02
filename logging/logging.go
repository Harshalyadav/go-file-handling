package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fileInfo, err := os.OpenFile("./1text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("Hello log")
	log.SetOutput(fileInfo)

}
