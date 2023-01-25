package bookstore

// Book represents information about a book
type Book struct {
	Id     int
	Title  string
	Author string
	Copies int
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
