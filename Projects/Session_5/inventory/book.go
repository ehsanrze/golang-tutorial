package inventory

// Book represents a book in the inventory
type Book struct {
	ID     string
	Name   string
	Author string
	Price  float64
}

// GetID returns the ID of the book
func (b Book) GetID() string {
	return b.ID
}

// GetName returns the name of the book
func (b Book) GetName() string {
	return b.Name
}

// GetPrice returns the price of the book
func (b Book) GetPrice() float64 {
	return b.Price
}

func (b *Book) ApplyDiscount(discount float64) {
	b.Price = b.Price * (1 - discount/100)
}
