package main

import "fmt"

func main() {
	book := "Sapiens"

	fmt.Println("Curent book:", book)
	updateBook(&book)
	fmt.Println("Post-update book:", book)
}

func updateBook(book *string) {
	*book = "Circe"
	fmt.Println("Updated book: ", *book)
	//return *book
}
