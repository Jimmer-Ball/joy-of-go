package bookstore

import (
	"errors"
	"fmt"
)

// Book represents information about a book
type Book struct {
	Id              int
	Title           string
	Author          string
	Copies          int
	PriceCents      int
	DiscountPercent int
}

// Equals means a book is equal to another book, if either both book pointers provided
// point to the same struct in memory, or if both structs have the same Id value, as
// Id is the Primary Key of the struct
func Equals(a, b *Book) bool {
	var returnValue bool
	if a == b {
		// Both a and b point to the same structure
		returnValue = true
	} else {
		// Unlike C++, you don't need to dereference the pointers in a function to "get" at
		// the structure pointed to
		if a.Id == b.Id {
			// Different structs with the same Id, means logical equality
			returnValue = true
		}
	}
	return returnValue
}

// DiscountPriceCents method equals the PriceCents minus whatever saving the DiscountPercent represents. This
// function has a "receiver" element that comes after the "func" part, that indicates what struct or
// interface the "method" can be applied to, and is what makes this function a "method". Here the receiver is
// "(b Book)" which is then used within the method. So it is GO's way of saying here is a method.  The received
// type can be an interface too, which makes it nicely adaptive
func (b *Book) DiscountPriceCents() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}

// AdjustPrices is a receiver method that takes a pointer to a book struct. As such the Book pointed to gets
// updated, meaning encapsulation.Since methods often need to modify their receiver, pointer receivers are
// often more common than value receivers.
func (b *Book) AdjustPrices(newPrice int, newDiscountPercent int) error {
	var err error = nil
	if newPrice > 0 && (newDiscountPercent > 0 && newDiscountPercent < 100) {
		b.PriceCents = newPrice
		b.DiscountPercent = newDiscountPercent
	} else {
		err = errors.New(fmt.Sprintf("The PriceCents value %d and DiscountPercent %d value are invalid",
			newPrice, newDiscountPercent))
	}
	return err
}
