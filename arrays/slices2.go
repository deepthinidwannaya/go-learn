package main

import "fmt"

func main() {
	movies := []string{
		"Grace and Frankie",
		"Fauda",
		"The Good Place",
		"Kim's Convenience",
		"WandaVision",
		"Mandalorian"}

	movies1 := movies[1:3]
	movies2 := movies[0:2]

	fmt.Println(movies1)
	fmt.Println(movies2)

	movies1[0] = "Kingdom"

	fmt.Println(movies1)
	fmt.Println(movies2)
	fmt.Println(movies)
}
