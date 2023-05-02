package main

import (
	"errors"
	"fmt"
)

func process(n int) error {
	if n%2 == 0 {
		return errors.New("Only Odd numbre is allowed")
	}
	return nil
}

func check(e error) {

	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("Operation sucessful")
}

func main() {
	e := process(2)
	check(e)
	e = process(3)
	check(e)

}
