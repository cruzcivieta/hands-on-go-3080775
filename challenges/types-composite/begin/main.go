// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title  string
	author author
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (library library) addBook(book book) {
	library[book.author.name] = append(library[book.author.name], book)
}

// define a lookupByAuthor function t
func (library library) lookupByAuthor(author author) []book {
	return library[author.name]
}

func main() {
	// create a new library
	library := make(library)

	// add 2 books to the library by the same author
	library.addBook(
		book{
			author: author{
				name: "J.R.R. Tolkien",
			},
			title: "The Lord of the rings - The Ring Goes South",
		})
	library.addBook(
		book{
			author: author{
				name: "J.R.R. Tolkien",
			},
			title: "The Lord of the rings - The Ring Sets Out",
		})

	// add 1 book to the library by a different author
	library.addBook(
		book{
			author: author{
				name: "Jeff Linsday",
			},
			title: "Dexter. The dark passenger",
		})

	// dump the library with spew
	fmt.Printf("%#v\n", library)

	// lookup books by known author in the library
	var books = library.lookupByAuthor(author{name: "Jeff Linsday"})
	var booksTolkien = library.lookupByAuthor(author{name: "J.R.R. Tolkien"})

	// print out the first book's title and its author's name
	fmt.Printf("%#v\n", books[0])
	fmt.Printf("%#v\n", booksTolkien[0])
}
