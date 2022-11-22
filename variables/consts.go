package main

import "fmt"

func main() {
	const c = 3.14
	fmt.Println("Constant is: ", c)
	//ERROR:
	//c = 1.342423
}
