package bookstore_test

import (
	"bookstore"
	"github.com/google/go-cmp/cmp"
	"testing"
)

// Buying a book reduces the number of copies available
func TestBuy(t *testing.T) {
	t.Parallel()
	purchase := bookstore.Book{Id: 1, Title: "Example Book", Author: "Dave Normal", Copies: 3}
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

func TestGetBook(t *testing.T) {
	t.Parallel()
	want := bookstore.Book{Id: 1, Title: "Example Book", Author: "Dave Normal", Copies: 3}
	got, _ := bookstore.GetBook(1)
	if want != got {
		t.Errorf("Wanted %d but got %d", want.Id, got.Id)
	}
}

func TestGetBookMissing(t *testing.T) {
	t.Parallel()
	_, err := bookstore.GetBook(122)
	if err == nil {
		t.Errorf("There is no book with Id of %d", 122)
	} else {
		t.Log("We got the expected error", err)
	}
}

// Getting all books returns a list of 3 books provided by the store
func TestGetAll(t *testing.T) {
	t.Parallel()
	// Expected
	want := []bookstore.Book{
		{Id: 1, Title: "Example Book", Author: "Dave Normal", Copies: 3},
		{Id: 2, Title: "For the Love of Go", Author: "John Arundel", Copies: 4},
		{Id: 3, Title: "Get me my gun", Author: "Raul Fandango"},
	}

	// got is an unlimited slice variable that can be used to reference an
	// underlying array of bookstore.Book items. See https://go.dev/doc/effective_go#slices
	var got []bookstore.Book
	got, _ = bookstore.GetAll()
	if got != nil {
		count := len(got)
		// Use the deep compare provided by the "cmp" module
		// instead of using the custom Equals
		if !cmp.Equal(want, got) {
			// want should be the same as got
			t.Errorf("Not what was wanted:\n %s", cmp.Diff(want, got))
		}
		if count == 3 {
			t.Log("We got 3 books like we wanted")

			// Slices are a reference into an existing collection of things
			if got[0].Title != "Example Book" &&
				got[1].Title != "For the Love of Go" &&
				got[2].Title != "Get me my gun" {
				t.Error("got books in wrong order")
			}

			// We can use a slice index to modify an array
			got[0].Title = "Updated Example Book"
			if got[0].Title != "Updated Example Book" {
				t.Error("Something is wrong here")
			}

			// We can add to an array using a slice
			got = append(got, bookstore.Book{Id: 4, Title: "Another book", Author: "Dave Normal", Copies: 2})
			if len(got) != 4 {
				t.Error("We appended to the got books listing, but the expected length is wrong")
			}
		} else {
			t.Errorf("We got the wrong number of books, size %d", count)
		}
	} else {
		t.Error("We expected some books")
	}
}

func TestEquals(t *testing.T) {
	t.Parallel()
	a := bookstore.Book{Id: 1, Title: "Example Book", Author: "Dave Normal", Copies: 3}
	b := bookstore.Book{Id: 1, Title: "Some Other Book", Author: "Steve Abnormal", Copies: 3}

	t.Logf("Pointer to \"a\" has value of %p", &a)
	t.Logf("Pointer to \"b\" has value of %p", &b)

	if !bookstore.Equals(&a, &b) {
		t.Error("Expected equality operator to be the same")
	}

	// d is a pointer to a
	var d = &a
	if !bookstore.Equals(d, &a) {
		t.Errorf("Expected %p pointer \"d\" to be the same as reference %p to \"&a\"", d, &a)
	} else {
		t.Logf("Pointer to \"a\" has value of %p", d)
		t.Logf("Id of \"a\" via pointer \"d\" is %d", d.Id)
		t.Logf("Title of \"a\" via pointer \"d\" is %s", d.Title)
	}
}
