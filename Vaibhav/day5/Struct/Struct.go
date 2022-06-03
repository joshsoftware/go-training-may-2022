package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func main() {
	var Book Books // Declare Book of type Books

	// book  specification
	Book.title = "Go Programming"
	Book.author = "Johnson"
	Book.subject = "Go Programming "
	Book.book_id = 6495407

	/* print Book1 info */
	printBook(&Book)

	// normal printing book ingo
	// fmt.Printf("Book  title : %s\n", Book.title)
	// fmt.Printf("Book  author : %s\n", Book.author)
	// fmt.Printf("Book  subject : %s\n", Book.subject)
	// fmt.Printf("Book  book_id : %d\n", Book.book_id)

}
