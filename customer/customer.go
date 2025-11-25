package customer

// Customer represents a customer with a name.
type Customer struct {
	Name string
}

// NewCustomer creates a new Customer instance with the given name.
func NewCustomer(name string) Customer {
	return Customer{Name: name}
}

// String returns customer's name as its string representation.
func (c Customer) String() string {
	return c.Name
}
