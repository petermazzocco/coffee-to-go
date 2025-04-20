package pointers

import (
	"fmt"
	"time"
)

type order struct {
	size    string
	variant string
}

// Dereference the pointer(*) to the order variable
// We need to make sure we're grabbing the values from our order
func changeOrderSize(o *order, newSize string) order {
	// We need to make sure we are changing the order we are pointing to
	// You can do that by using the dereference operator(*)
	*o = order{
		size:    newSize,      // change the order size here
		variant: (*o).variant, // we'll keep it as iced coffee
	}
	return *o
}

func UpdateOrder() {
	// Here is the original order we placed
	// We create a new struct
	order := order{
		size:    "medium",
		variant: "iced coffee",
	}

	fmt.Println("Our original order:", order)

	time.Sleep(1 * time.Second) // simulate order update request
	fmt.Println("Updating our order...")

	// We are pointing(&) to the order variable
	// We have to make sure we are updating our specific order
	newOrder := changeOrderSize(&order, "large") // run the change order function with our desired size

	time.Sleep(1 * time.Second) // simulate order update response
	fmt.Printf("Our updated order vs our original order: %v vs %v. Notice how it's the same now?\n", newOrder, order)
}
