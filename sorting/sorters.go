package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"test", "what", "how", "to", "get", "away", "with", "murder"}
	sort.Strings(names)
	fmt.Println(names)

	people := []struct {
		name string
		age  int
	}{
		{"Annalise", 40},
		{"trophy boy", 18},
		{"woof woof", 2},
	}

	sort.Slice(people,
		func(i, j int) bool {
			return people[i].age < people[j].age
		})
	fmt.Println(people)

	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 5
	idx := sort.Search(len(a), func(i int) bool {
		fmt.Println(a[i])
		return a[i] >= x
	})
	if idx < len(a) && a[idx] == x {
		fmt.Printf("Found %v\n", a[idx])
	} else {
		fmt.Printf("Insert it here %v\n", idx)
	}
}
