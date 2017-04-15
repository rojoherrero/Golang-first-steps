package models

// Product represents any type pf article a shop can have
type Product struct {
	name          string
	itemID        int
	cost          float32
	isAvailable   bool
	inventoryLeft int
}
