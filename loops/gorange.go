package main

import "fmt"

func main() {
	names := []string{
		"Thor",
		"Wanda",
		"Black Widow",
		"Star lord"}

	for index, name := range names {
		fmt.Println(name)
		fmt.Println(index)
	}
}
