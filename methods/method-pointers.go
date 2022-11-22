package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func (v *Vertex) increment() {
	// methods with pointer receivers take either a value or a pointer as the receiver when they are called
	v.x += 10
	v.y += 10
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.increment()
	fmt.Println(v)
}
