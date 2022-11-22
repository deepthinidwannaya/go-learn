package main

import (
	"fmt"
	"math"
)

func main() {
	if even, odd := 2, 3; even > odd {
		fmt.Println("Even is bigger")
	} else if math.Sqrt(4) == float64(even) {
		fmt.Println("Squared yay")
	} else {
		fmt.Println("Elsa")
	}
}
