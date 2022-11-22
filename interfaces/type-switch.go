package main

import "fmt"

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v is int\n", v)
	case float64:
		fmt.Printf("%v is float\n", v)
	case string:
		fmt.Printf("%v is string\n", v)
	}
}

func main() {
	checkType(1)
	checkType(1.0)
	checkType("test")
}
