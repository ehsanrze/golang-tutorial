// inventory/electronics.go
package inventory

// Electronics represents an electronic item in the inventory
type Electronics struct {
	ID       string
	Name     string
	Brand    string
	Price    float64
	Warranty int // in months
}

// GetID returns the ID of the electronic item
func (e *Electronics) GetID() string {
	return e.ID
}

// GetName returns the name of the electronic item
func (e *Electronics) GetName() string {
	return e.Name
}

// GetPrice returns the price of the electronic item
func (e *Electronics) GetPrice() float64 {
	return e.Price
}

func (e *Electronics) ApplyDiscount(discount float64) {
	e.Price = e.Price * (1 - discount/100)
}
