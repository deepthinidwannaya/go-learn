package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	// Use camel case
	animalName string
	species    string
	category   int
)

//
//var (
//	course, name string
//	module, clip int
//)

var (
	course, name = "Test", "name"
	module, clip = 1, "40"
)

func main() {
	fmt.Println("Name and course are set to ", name, " and", course)
	fmt.Println("Module and clip are set to ", module, " and", clip)
	fmt.Println("Name is of type: ", reflect.TypeOf(name))
	fmt.Println("category is of type: ", reflect.TypeOf(category))
	iClip, err := strconv.Atoi(clip)

	if err == nil {
		total := module + iClip
		fmt.Println(total)
	}
	fmt.Println("Memory address of clip is :", &clip)

	var ptr = &clip
	//*ptr="3432"
	fmt.Println("Pointing clip variable at address,", ptr, " with value", *ptr)
}
