package main

import "fmt"

type Book struct {
	Title         string
	ID            int
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title: \t\t%q\n"+
			"Author: \t\t%q\n"+
			"Year: \t\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	Book{
		ID:            1,
		Title:         "Hitchhiker's Guide To The Galaxy",
		Author:        "Douglas Adams",
		YearPublished: 1979,
	},
	Book{
		ID:            2,
		Title:         "Beautiful World, Where Are You",
		Author:        "Sally Rooney",
		YearPublished: 2021,
	},
	Book{
		ID:            3,
		Title:         "Sapiens",
		Author:        "Yuval Noah Harari",
		YearPublished: 2018,
	},
}
