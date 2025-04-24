package interfaces

import "fmt"

// Here we have two importers that carry the same
// functions, except there's one difference.
// Because of these insane tariffs, one of our
// countries is being tariffed, we need to pay a fee

// We will need to buy supple from both of our importers
// Meaning we can set an interface to handle that with
// minimal code. But the interface will be used to handle the
// tariff fee as well for ONE of our importers
type Importer interface {
	BuySupply(amount int)
}

// Define our Importer structs
type InternationalImporter struct {
	price int
}
type LocalImporter struct {
	price int
}

// Constructor for new importers
func NewInternationalImporter() *InternationalImporter {
	return &InternationalImporter{
		price: 200,
	}
}
func NewLocalImporter() *LocalImporter {
	return &LocalImporter{
		price: 1000,
	}
}

// Our buy supply functions for each unique importer
func (i *InternationalImporter) BuySupply(orderAmount int) {
	amount := orderAmount*i.price + 200 // 200 represents the tariff fee
	fmt.Printf("We are paying $%d to get from our international importer for %d orders\n", amount, orderAmount)
}
func (l *LocalImporter) BuySupply(orderAmount int) {
	amount := orderAmount*l.price + 0 // 0 represents no tariff fee
	fmt.Printf("We are paying $%d to get from our local importer for %d orders\n", amount, orderAmount)
}

func CoffeeShopChangeImport() {
	// Create new importers
	importers := []Importer{
		NewInternationalImporter(),
		NewLocalImporter(),
	}

	for _, importer := range importers {
		importer.BuySupply(10) // Let's buy supply from each
	}
}
