package main

import "fmt"

type Shape interface {
	Area() float64
	Parameter() float64
}

type Rectangle struct {
	length, width float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Parameter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	var s Shape = Rectangle{4, 6}
	fmt.Println(s.Area())
	fmt.Println(s.Parameter())

}