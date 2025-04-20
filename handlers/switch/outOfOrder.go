package switchpkg

import (
	"fmt"
)

func CoffeeShopOutOfOrder(variant string) {
	menu := map[string]bool{
		"Coffee":   true,
		"Tea":      true,
		"Espresso": false,
		"Latte":    true,
		"Mocha":    false,
		"Water":    false,
	}
	// Let's verify the item exists on our menu
	_, exists := menu[variant]
	if !exists {
		fmt.Println("Invalid variant. Please see the menu for our variants!")
		return
	}
	// If the variant returns an item that is sold out, we will let them know.
	switch variant {
	case "Coffee":
		fmt.Println("Coffee is out of order. Please try another item.")
	case "Tea":
		fmt.Println("Tea is out of order. Please try another item.")
	case "Espresso":
		fmt.Println("Espresso is out of order. Please try another item.")
	case "Latte":
		fmt.Println("Latte is out of order. Please try another item.")
	case "Water":
		fmt.Println("Water is out of order. Please try another item.")
	default:
		fmt.Println("Invalid variant. Please see the menu for our variants!")
	}
}
