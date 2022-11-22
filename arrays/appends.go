package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2}
	b := append(a, 3)
	fmt.Printf("cap(a)=%v,cap(b)=%v\n", cap(a), cap(b))
	c := append(b, 4)
	d := append(b, 5)
	fmt.Printf("cap(a)=%v,cap(b)=%v\n", cap(a), cap(b))
	fmt.Println(a, b, c, d)
	fmt.Printf("%p,%p,%p,%p\n", a, b, c, d)
}
