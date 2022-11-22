package main

import "fmt"

type Square struct {
	side float64
}

// Receiver argument - type defined between func keyword and method name
func (s Square) Area() float64 {
	return s.side * s.side
}

func main() {
	s := Square{5}
	fmt.Println(s.Area())
}
