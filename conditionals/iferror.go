package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("fsdsf")

	if err != nil {
		fmt.Println("Oops")
	}
}
