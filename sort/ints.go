package main

import (
	"fmt"
	"sort"
)

func main() {
	vars := []int{1, 7, 3, 4, 2, 7, 8}
	sort.Ints(vars)
	fmt.Println(vars)
}