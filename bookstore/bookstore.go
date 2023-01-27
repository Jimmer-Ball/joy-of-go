package bookstore

// Note that uppercase methods and structs are visible outside this package
import (
	"fmt"
	"sort"
)

// Rating represent the number of Likes and Dislikes registered against a given Book
type Rating struct {
	Id       int
	Likes    int
	Dislikes int
}

// The stock catalog needs to be a common shared resource
var stock = []Book{
	{Id: 1, Title: "Example Book", Author: "Dave Normal", Copies: 3},
	{Id: 2, Title: "For the Love of Go", Author: "John Arundel", Copies: 4},
	{Id: 3, Title: "Get me my gun", Author: "Raul Fandango"},
}

// ratings register Likes and Dislikes against a book
var ratings = map[int]Rating{}

// Like increments the like count against a Book
func Like(Id int) error {
	return ratingsManagement(Id, true)
}

// Dislike increments the dislike count against a Book
func Dislike(Id int) error {
	return ratingsManagement(Id, false)
}

// GetRating provides the Rating for a given book
func GetRating(Id int) (Rating, error) {
	var err error = nil
	var rating Rating
	var ok bool

	_, err = GetBook(Id)
	if err == nil {
		rating, ok = ratings[Id]
		if !ok {
			err = fmt.Errorf("rating for book %d does not exist", Id)
		}
	}
	return rating, err
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
			err = fmt.Errorf("no copies of book %d are available", Id)
		} else {
			// A book can only be purchased if stock is available
			if gotBook.Copies >= copies {
				// Reduce stock
				updatedBook = DecreaseStock(Id, copies)
			} else {
				err = fmt.Errorf("cannot buy %d copies of book %d, as only %d are available",
					copies, Id, gotBook.Copies)
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
		err = fmt.Errorf("the book with Id %d does not exist", Id)
	}
	return gotBook, err
}

// GetAll books returns details of all books
func GetAll() []Book {
	// Apply a function literal to do the slice sorting consistently
	sort.Slice(stock, func(i, j int) bool {
		return stock[i].Id < stock[j].Id
	})
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

func ratingsManagement(Id int, like bool) error {
	var err error = nil

	_, err = GetBook(Id)
	if err == nil {
		// So here we get a copy of the underlying entry in the map,
		// and when we are done, we update the map with the copy.
		// This is a way of "avoiding" pointers and de-referencing
		rating, ok := ratings[Id]
		if ok {
			if like {
				rating.Likes = rating.Likes + 1
			} else {
				rating.Dislikes = rating.Dislikes + 1
			}
		} else {
			rating.Id = Id
			if like {
				rating.Likes = 1
			} else {
				rating.Dislikes = 1
			}
		}
		ratings[Id] = rating
	}
	return err
}
