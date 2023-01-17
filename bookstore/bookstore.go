package bookstore

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
