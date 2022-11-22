package main

import "fmt"

func main() {
	best := champions(1, 2, 3, 4, 5)
	fmt.Println(best)
}

func champions(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i > best {
			best = i
		}
	}
	return best
}
