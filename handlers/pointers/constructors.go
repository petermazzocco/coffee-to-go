package pointers

import "fmt"

type CoffeeShop struct {
	name    string
	address string
	hours   int
}

// In order for us to create a new coffee shop via the
// constructor, we must pass in a dereferenced CoffeeShop struct pointer
// as an argument.
// This is because the constructor function takes a dereferenced
// pointer to a CoffeeShop struct so that within the function, we will
// point to the CoffeeShop struct we are creating and then return the
// dereferenced pointer to it.
func NewCoffeeShop(c *CoffeeShop) *CoffeeShop {
	return &CoffeeShop{
		name:    c.name,
		address: c.address,
		hours:   c.hours,
	}
}

func CreateNewCoffeeShop() {
	// We pass in the pointer to the CoffeeShop struct as the argument
	// in which the NewCoffeeShop constructor will deference and update
	coffeeShop := NewCoffeeShop(&CoffeeShop{
		name:    "My Coffee Shop",
		address: "123 Main St",
		hours:   8,
	})

	fmt.Println("You created a new coffee shop!", *coffeeShop)
}
