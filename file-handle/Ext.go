package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	r := filepath.Join("hello","Hi","I'm harshal")
	path := filepath.Join("dir1","dir2/../dir3","text.txt")
	fmt.Println(r)
	fmt.Println(filepath.Ext(path))
	fmt.Println(filepath.IsAbs("/dir/file"))
}