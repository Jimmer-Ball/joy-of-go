package bookstore

// Note that uppercase methods and structs are visible outside this package
import (
	"errors"
	"fmt"
)

// Buy a book decrements the number of copies if found by the number of copies expected and
// returns either an error or returns the updated book details in the catalog
func Buy(b Book) (Book, error) {

	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func GetBook(Id int) (Book, error) {
	var gotBook Book
	var err error = nil

	// Look for the book in the catalog
	var books []Book
	books, _ = getCatalog()
	for _, b := range books {
		if b.Id == Id {
			gotBook = b
		}
	}

	// If the book was not found (where the field has default value of zero), report an error
	if gotBook.Id == 0 {
		errorMessage := fmt.Sprintf("The book with Id %d does not exist", Id)
		err = errors.New(errorMessage)
	}

	return gotBook, err
}

// GetAll books returns details of all books
func GetAll() ([]Book, error) {
	return getCatalog()
}

func getCatalog() ([]Book, error) {
	return []Book{
		{Id: 1, Title: "Example Book", Author: "Dave Normal", Copies: 3},
		{Id: 2, Title: "For the Love of Go", Author: "John Arundel", Copies: 4},
		{Id: 3, Title: "Get me my gun", Author: "Raul Fandango"},
	}, nil
}
