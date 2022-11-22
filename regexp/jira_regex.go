package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "CSCM-1435345+"
	str2 := "test"

	re := regexp.MustCompile(`^[A-Za-z0-9\-]+$`)
	fmt.Printf("Pattern: %v\n", re.String())      // print pattern
	fmt.Println("Matched:", re.MatchString(str1)) // true
	fmt.Println("Matched:", re.MatchString(str2)) // true
}
