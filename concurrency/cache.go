package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	for i := 0; i < 10; i++ {
		id := rnd.Intn(3) + 1
		if b, ok := queryCache(id); ok {
			fmt.Println("From cache")
			fmt.Println(b)
			continue
		}
		if b, ok := queryDb(id); ok {
			fmt.Println("From DB")
			fmt.Println(b)
			continue
		}
		fmt.Printf("Book not found with id %v\n", id)
		time.Sleep(150 * time.Millisecond)
	}

}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}
func queryDb(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			return b, true
		}
	}
	return Book{}, false
}
