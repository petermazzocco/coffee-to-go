package generics

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Create an generic Item struct to track the name and price
// The price will be any comparable integer type
type Item[T comparable, V constraints.Integer] struct {
	name  string
	price V
}

type CustomInventory[T comparable, V Item[T, int]] map[T]V

// Create a map to store the items that properly implements the Item struct
var newItems = map[int]Item[int, int]{
	1: {name: "Coffee", price: 10},
	2: {name: "Espresso", price: 1},
	3: {name: "Cappuccino", price: 20},
}

func NewInventory() {
	// Loop through out inventory
	inventory := make(CustomInventory[int, Item[int, int]])
	for i, v := range newItems {
		inventory[i] = Item[int, int]{name: v.name, price: v.price}
	}
	fmt.Println("Here's our inventory", inventory)
}
