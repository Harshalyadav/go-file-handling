package main

import (
	"fmt"
	"sort"
)

func main() {
	vars := []string{"Learning","Go","Language"}
	sort.Strings(vars)
	fmt.Println(vars)
}