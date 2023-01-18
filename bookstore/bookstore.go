package bookstore

// Note that uppercase methods and structs are visible outside this package
import "errors"

// Customer represents information about a customer.
type Customer struct {
	Id    string
	Name  string
	Email string
}

// Book represents information about a book
type Book struct {
	Id     string
	Title  string
	Author string
	Copies int
}

// Buy a book decrements the number of copies. Weirdly, multiple return statements is fine in GO
func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

// GetAll books returns details of all books
func GetAll() ([]Book, error) {
	return []Book{
		{Title: "Example Book", Author: "Dave Normal", Copies: 3},
		{Title: "For the Love of Go", Author: "John Arundel", Copies: 4},
		{Title: "Get me my gun", Author: "Raul Fandango"},
	}, nil
}
