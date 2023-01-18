package bookstore_test

import (
	"bookstore"
	"testing"
)

// Buying a book reduces the number of copies available
func TestBuy(t *testing.T) {
	t.Parallel()
	purchase := bookstore.Book{Title: "Example Book", Author: "Dave Normal", Copies: 3}
	want := 2
	result, _ := bookstore.Buy(purchase)
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

// Buying a book when there are none in stock results in an error
func TestBuyNoneAvailable(t *testing.T) {
	t.Parallel()
	purchase := bookstore.Book{Title: "Missing Book", Author: "Gone Missing", Copies: 0}
	_, err := bookstore.Buy(purchase)
	if err != nil {
		// We expected an error
		t.Log("Expected error returned:", err)
	} else {
		t.Error("We expected an error, as you cannot buy a book if there are none left")
	}
}

// Getting all books returns a list of 3 books provided by the store
func TestGetAll(t *testing.T) {
	t.Parallel()
	// books is an unlimited slice variable that can be used to reference an
	// underlying array of bookstore.Book items. See https://go.dev/doc/effective_go#slices
	var books []bookstore.Book
	books, _ = bookstore.GetAll()
	if books != nil {
		count := len(books)
		if count == 3 {
			t.Log("We got 3 books like we expected")

			// Slices are a reference into an existing collection of things
			if books[0].Title != "Example Book" &&
				books[1].Title != "For the Love of Go" &&
				books[2].Title != "Get me my gun" {
				t.Error("Books not in the right order")
			}

			// We can use a slice index to modify an array
			books[0].Title = "Updated Example Book"
			if books[0].Title != "Updated Example Book" {
				t.Error("Something is wrong here")
			}

			// We can add to an array using a slice
			books = append(books, bookstore.Book{Title: "Another book", Author: "Dave Normal", Copies: 2})
			if len(books) != 4 {
				t.Error("We appended to the books listing, but the expected length is wrong")
			}
		} else {
			t.Errorf("We got the wrong number of books, size %d", count)
		}
	} else {
		t.Error("We expected some books")
	}
}
