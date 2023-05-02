package main

import "os"

func main() {

	file, err := os.OpenFile("./text.txt",os.O_APPEND| os.O_CREATE|os.O_WRONLY,0060)
}