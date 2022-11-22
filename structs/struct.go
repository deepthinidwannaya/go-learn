package main

import "fmt"

func main() {
	type movieMeta struct {
		movieName string
		rating    float64
		platform  string
	}

	//movie1 := new(movieMeta)

	movie1 := movieMeta{
		"Venom",
		4.2,
		"Netflix"}
	fmt.Println(movie1)
}
