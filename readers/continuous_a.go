package main

import "fmt"

/*
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/

type MyReader struct{}

func (mr MyReader) Read(b []byte) (int, error) {
	return copy(b, "A"), nil
}

func main() {
	b := make([]byte, 8)
	r := MyReader{}
	for {
		n, err := r.Read(b)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("%v %v", n, b)
	}
}
