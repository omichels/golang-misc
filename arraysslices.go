package main

import "fmt"

type Book struct {
	id    int
	title string
}

func (b Book) String() string {
	return fmt.Sprintf("Book id: %d, title: %s", b.id, b.title)
}

func main() {

	var a [4]int
	a[0] = 1

	fmt.Println(a)
	fmt.Println(a[0])

	var books []Book

	books = append(books, Book{
		id:    1,
		title: "of men and mice",
	})

	fmt.Println(books[0])

}
