package main

import (
	"fmt"
	"time"
)

func speak(word string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(word)
	}
}
func main() {
	go speak("world")
	speak("hello")
}
