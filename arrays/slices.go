package main

import "fmt"

func main() {
	// slice creation
	courses := make([]string, 5, 10)
	courses[0] = "A"
	courses[1] = "B"
	courses[2] = "C"
	courses[3] = "C"
	courses[4] = "C"
	// Extending a slice's capacity
	courses = courses[:6]
	courses[5] = "CS"

	// courses = courses[:12] slice bounds out of range

	fmt.Println(len(courses), cap(courses))

	fmt.Println(courses)
}
