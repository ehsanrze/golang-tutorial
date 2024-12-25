package main

import (
	"fmt"
)

// Book struct
type Book struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	IsBorrowed bool   `json:"is_borrowed"`
}

// NewBook Constructor for Book
func NewBook(title, author string) Book {
	return Book{Title: title, Author: author, IsBorrowed: false}
}

// Info Book method to display information
func (b Book) Info() string {
	return fmt.Sprintf("Title: %s, Author: %s, Borrowed: %t", b.Title, b.Author, b.IsBorrowed)
}

// Library struct
type Library struct {
	Books []Book
}

// AddBook Add a book to the library
func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

// ListAvailableBooks List available books
func (l Library) ListAvailableBooks() {
	for _, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Println(book.Info())
		}
	}
}

// BorrowBook Borrow a book
func (l *Library) BorrowBook(title string) error {
	for i, book := range l.Books {
		if book.Title == title {
			if book.IsBorrowed {
				return fmt.Errorf("book '%s' is already borrowed", title)
			}
			l.Books[i].IsBorrowed = true
			return nil
		}
	}
	return fmt.Errorf("book '%s' not found", title)
}

// ReturnBook Return a book
func (l *Library) ReturnBook(title string) error {
	for i, book := range l.Books {
		if book.Title == title {
			if !book.IsBorrowed {
				return fmt.Errorf("book '%s' was not borrowed", title)
			}
			l.Books[i].IsBorrowed = false
			return nil
		}
	}
	return fmt.Errorf("book '%s' not found", title)
}

// Main function
func main() {
	lib := Library{}

	lib.AddBook(NewBook("1984", "George Orwell"))
	lib.AddBook(NewBook("The Great Gatsby", "F. Scott Fitzgerald"))

	fmt.Println("Available books:")
	lib.ListAvailableBooks()

	fmt.Println("\nBorrowing '1984'...")
	err := lib.BorrowBook("1984")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nAvailable books after borrowing:")
	lib.ListAvailableBooks()

	fmt.Println("\nReturning '1984'...")
	err = lib.ReturnBook("1984")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nAvailable books after returning:")
	lib.ListAvailableBooks()
}
