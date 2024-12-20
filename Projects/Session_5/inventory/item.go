package inventory

// Item defines the interface for inventory items
type Item interface {
	GetID() string
	GetName() string
	GetPrice() float64
	ApplyDiscount(discount float64)
}
