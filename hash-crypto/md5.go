package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func endcodeWithMD5(pass string) string {
	var hashCode = md5.Sum([]byte(pass))
	return hex.EncodeToString(hashCode[:])
}

func main() {

	var password string
	fmt.Println("Enter a password")
	fmt.Scanln(&password)
	fmt.Println("Decrypt password is", endcodeWithMD5(password))
}
