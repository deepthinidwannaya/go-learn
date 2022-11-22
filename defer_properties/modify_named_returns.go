package main

import "fmt"

func main() {
	//Deferred functions may read and assign to the returning function’s named return values.
	fmt.Println(c())

}

func c() (i int) {
	defer func() {
		i++
	}()
	return 10
}
