package maps

import "fmt"

type Item struct {
	soldOut bool
	price   float64
}

// Our available menu!
var menu = map[string]Item{
	"Espresso":   {soldOut: false, price: 2.99},
	"Latte":      {soldOut: false, price: 3.99},
	"Cappuccino": {soldOut: true, price: 3.99},
	"Americano":  {soldOut: false, price: 2.99},
	"Mocha":      {soldOut: true, price: 3.99},
	"Tea":        {soldOut: false, price: 1.99},
	"Water":      {soldOut: false, price: 1.99},
}

// Set a global variable for us to access in the handlers
var Menu = menu

func AddNewMenuItem(name string, price float64) {
	menu[name] = Item{soldOut: false, price: price}
	fmt.Println("Woohoo! We've added a new item to our menu:", name)
	fmt.Println("Now our menu has", len(menu), "items!")
	for item, details := range menu {
		fmt.Printf(">>> 🥤%s - $%.2f - sold out: %t\n", item, details.price, details.soldOut)
	}
}

func RemoveMenuItem(name string) {
	// Let's verify we even have the item on our menu:
	_, ok := menu[name]
	if !ok {
		fmt.Println("The item you are trying to remove does not exist on our menu! Run the `coffee order --menu` to view our menu")
		return
	}
	delete(menu, name)
	fmt.Printf("Looks like customers weren't buying our %s anyways\n", name)
	fmt.Println("Now our menu has", len(menu), "items!")
	for item, details := range menu {
		fmt.Printf(">>> 🥤%s - $%.2f - sold out: %t\n", item, details.price, details.soldOut)
	}
}

func CoffeeShopMenu() {
	fmt.Println("The Coffee Shop Menu:")
	for item, details := range menu {
		fmt.Printf(">>> 🥤%s - $%.2f - sold out: %t\n", item, details.price, details.soldOut)
	}

	fmt.Println("Would you like to add a new menu item? If so, run the cmd `coffee order --new [item-name] [price]`")
}
