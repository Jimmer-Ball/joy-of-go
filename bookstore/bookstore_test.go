package bookstore_test

import (
	"bookstore"
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

// Liking a book increments a book's like count
func TestLike(t *testing.T) {

	_ = bookstore.Like(1)
	_ = bookstore.Like(1)
	_ = bookstore.Like(1)
	_ = bookstore.Like(2)

	var likeCount int
	likeCount, _ = bookstore.GetLikes(1)
	if likeCount != 3 {
		t.Errorf("We should have three likes against book Id 1")
	}
	likeCount, _ = bookstore.GetLikes(2)
	if likeCount != 1 {
		t.Errorf("We should have 1 like against book Id 2")
	}
	var err error = nil
	likeCount, err = bookstore.GetLikes(4)
	if likeCount != 0 {
		t.Errorf("We should have 0 like against book Id 4")
	}
	if err == nil {
		t.Error("We should have an error as book 4 does not exist")
	} else {
		if err.Error() != "the book with Id 4 does not exist" {
			t.Errorf("Wrong error message provided %s", err.Error())
		}
	}
}

// Buying a book reduces the number of copies available
func TestBuyAvailable(t *testing.T) {
	t.Parallel()

	// Get hold of a book
	targetBook := bookstore.Book{Id: 1, Title: "Example Book", Author: "Dave Normal", Copies: 3}
	got, _ := bookstore.GetBook(1)
	if targetBook != got {
		t.Errorf("Wanted %d but got %d", targetBook.Id, got.Id)
	}

	// Buy one copy and check stock reduction
	stockAfterPurchase, _ := bookstore.Buy(1, 1)
	if stockAfterPurchase.Copies != 2 {
		t.Errorf("The number of copies should have gone down by one")
	}

	// Re-get it and check stock level after purchase
	reGot, _ := bookstore.GetBook(1)
	if reGot.Copies != 2 {
		t.Error("The number of copies in stock should be 2")
	}

	bookstore.IncreaseStock(1, 1)

	// Re-get it and check stock level after restock
	againGot, _ := bookstore.GetBook(1)
	if againGot.Copies != 3 {
		t.Error("The number of copies in stock should be 3")
	}
}

// Buying a book with an unknown id results in an error
func TestBuyNotFound(t *testing.T) {
	t.Parallel()
	_, err := bookstore.Buy(99, 1)
	if err != nil {
		// We expected a Not in Stock error
		if err.Error() != "the book with Id 99 does not exist" {
			t.Errorf("Wrong error message provided %s", err.Error())
		}
	} else {
		t.Error("We expected an error, as you cannot buy a book if there are none left")
	}
}

// Buying too many copies of a book given the current stock levels results in an error
func TestBuyUnavailable(t *testing.T) {
	t.Parallel()
	_, err := bookstore.Buy(1, 12)
	if err != nil {
		// We expected a Cannot buy error message
		if !strings.Contains(err.Error(), "cannot buy 12 copies of book 1") {
			t.Errorf("Wrong error message provided %s", err.Error())
		}
	} else {
		t.Error("We expected an error, as you cannot buy a book if there are none left")
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	want := bookstore.Book{Id: 1, Title: "Example Book", Author: "Dave Normal", Copies: 3}
	got, _ := bookstore.GetBook(1)
	if want.Id != got.Id {
		t.Errorf("Wanted %d but got %d", want.Id, got.Id)
	}
}

func TestGetBookMissing(t *testing.T) {
	t.Parallel()
	_, err := bookstore.GetBook(122)
	if err == nil {
		t.Errorf("We should have an error, as there is no book with Id of %d", 122)
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
	var got = bookstore.GetAll()
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
