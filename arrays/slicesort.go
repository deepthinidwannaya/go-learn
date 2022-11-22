package main

import (
	"fmt"
	"sort"
)

func main() {
	//slice := make([]int, 5, 10)
	arr := []int{5, 8, 2, 6, 3}
	slice := arr[:3]
	fmt.Println(slice)
	sort.Ints(arr)
	fmt.Println(slice)
	fmt.Println(arr)

}
