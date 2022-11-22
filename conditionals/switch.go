package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	basicSwitch()
	fallthroughSwitch()
	idiomaticSwitch()
	switchWithoutConditions()
}

func basicSwitch() {
	switch "Test" {
	case "egg":
		fmt.Println("Egg!")
	case "test":
		fmt.Println("test")
	case "nest":
		fmt.Println("net")
	default:
		fmt.Println("default")
	}
}

func fallthroughSwitch() {
	switch "egg" {
	case "egg":
		fmt.Println("Egg!")
		fallthrough
	case "test":
		fmt.Println("test")
	case "nest":
		fmt.Println("net")
	default:
		fmt.Println("default")
	}
}

func idiomaticSwitch() {
	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("Its even")
	case 1, 3, 5, 7, 9:
		fmt.Println("It's odd")

	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}

/*
Switch without a condition is the same as switch true.

This construct can be a clean way to write long if-then-else chains.
Go's switch cases need not be constants, and the values involved need not be integers.
*/
func switchWithoutConditions() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
