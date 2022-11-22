package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	_, ok := i.(float64)
	fmt.Println(ok)
}
