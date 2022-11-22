package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// Empty return but named in definition
	return
}

func main() {
	fmt.Println(split(17))
	fmt.Println(1_5)
}
