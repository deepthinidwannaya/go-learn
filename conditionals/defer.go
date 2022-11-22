package main

import "fmt"

func another() {
	defer fmt.Println("test")
	fmt.Println("In another")
}

func main() {
	defer fmt.Println("world")
	another()
	fmt.Println("hello")
}
