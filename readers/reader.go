package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("%v %v", n, err)
		fmt.Printf("%q", b[:n])
		if err == io.EOF {
			break
		}
	}
}
