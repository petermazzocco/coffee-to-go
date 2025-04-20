package interfaces

import "fmt"

// Here we have two types of coffee we sell
// We want to update their origin and roast
// thanks to the tariffs
type coffee struct {
	variant string
	origin  string
	roast   string
}

// We define an interface for updating coffee data
// This will allow us to easily switch between
// different sources of coffee data
type Importer interface {
	SetOrigin(string) string
	SetRoast(string) string
}

// We can now update our origin and roast of a specific coffee
// We will dereference the pointer (*) to modify the specific coffee struct
func (c *coffee) SetOrigin(newOrigin string) string {
	c.origin = newOrigin
	return c.origin
}
func (c *coffee) SetRoast(newRoast string) string {
	c.roast = newRoast
	return c.roast
}

func CoffeeShopChangeImport() {
	// Let's make some of our coffees and their original origins and roasts
	// Feel free to add more variants, origins, and roasts
	// Notice that we are using the make() for this. It may be overkill
	// but it's good to practice
	coffees := make([]coffee, 2)              // make an array of 2 coffee structs
	variants := []string{"Espresso", "Latte"} // make an array of 2 variants
	origins := []string{"Colombia", "Kenya"}  // make an array of 2 origins
	roasts := []string{"Light", "Medium"}     // make an array of 2 roasts
	for i := range coffees {
		coffees[i] = coffee{
			variant: variants[i],
			origin:  origins[i],
			roast:   roasts[i],
		}
	}

	// Let's make sure we're importing the correct original data
	for i := range coffees {
		fmt.Printf("Our original %s comes from %s, and has a roast of %s\n", coffees[i].variant, coffees[i].origin, coffees[i].roast)
	}

	// Update the origin and roast of each of our coffees
	// We use the pointer (&) to modify the coffee struct
	newOrigins := []string{"Brazil", "Africa"} // You can change the new location here
	newRoasts := []string{"Light", "Medium"}   // You can change the new roast here
	for i := range coffees {
		(&coffees[i]).SetOrigin(newOrigins[i])
		(&coffees[i]).SetRoast(newRoasts[i])
	}

	// Print the updated coffee data
	for i := range coffees {
		fmt.Printf("Our updated %s now comes from %s, and has a roast of %s\n", coffees[i].variant, coffees[i].origin, coffees[i].roast)
	}
}
