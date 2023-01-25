package bookstore

// Note that uppercase methods and structs are visible outside this package
import (
	"errors"
	"fmt"
)

// The stock catalog needs to be a common shared resource
var stock = []Book{
	{Id: 1, Title: "Example Book", Author: "Dave Normal", Copies: 3},
	{Id: 2, Title: "For the Love of Go", Author: "John Arundel", Copies: 4},
	{Id: 3, Title: "Get me my gun", Author: "Raul Fandango"},
}

// Buy a book decrements the number of copies of a book in the stock catalog by the number of copies expected and
// returns either an error or the updated book details in the stock catalog
func Buy(Id int, copies int) (Book, error) {
	var gotBook Book
	var updatedBook Book
	var err error = nil

	gotBook, err = GetBook(Id)
	if err == nil {
		if gotBook.Copies == 0 {
			errorMessage := fmt.Sprintf("No copies of book %d are available", Id)
			err = errors.New(errorMessage)
		} else {
			// A book can only be purchased if stock is available
			if gotBook.Copies >= copies {
				// Reduce stock
				updatedBook = DecreaseStock(Id, copies)
			} else {
				errorMessage := fmt.Sprintf("Cannot buy %d copies of book %d, as only %d are available",
					copies, Id, gotBook.Copies)
				err = errors.New(errorMessage)
			}
		}
	}

	return updatedBook, err
}

// GetBook returns a copy of the book details for a book with a given id, or an error if no book with the id
// provided is found in stock
func GetBook(Id int) (Book, error) {
	var gotBook Book
	var err error = nil

	// Look for the book in the stock catalog
	for _, b := range stock {
		if b.Id == Id {
			gotBook = b
			break
		}
	}

	// If the book was not found in the slice, the value
	// of the id will still be default zero, not found
	if gotBook.Id == 0 {
		errorMessage := fmt.Sprintf("The book with Id %d does not exist", Id)
		err = errors.New(errorMessage)
	}

	return gotBook, err
}

// GetAll books returns details of all books
func GetAll() []Book {
	return stock
}

// IncreaseStock adds to the stock count for a given book in the stock catalog
func IncreaseStock(Id int, copies int) Book {
	return stockManagement(Id, copies, true)
}

// DecreaseStock reduces the stock count for a given book in the stock catalog
func DecreaseStock(Id int, copies int) Book {
	return stockManagement(Id, copies, false)
}

func stockManagement(Id int, copies int, increase bool) Book {
	var updatedBook Book
	for i, b := range stock {
		if b.Id == Id {
			// Reduce the book count held on the underlying stock catalog item, not in a copy of
			// a book in the catalog
			pr := &stock[i]
			if increase {
				(*pr).Copies = b.Copies + copies
			} else {
				(*pr).Copies = b.Copies - copies
			}
			// Return the updated book details by de-referencing the pointer
			updatedBook = *pr
			break
		}
	}
	return updatedBook
}
