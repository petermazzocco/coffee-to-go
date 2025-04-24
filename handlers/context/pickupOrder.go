package context

import (
	"context"
	"fmt"
)

// With context, you can pass state information to goroutines.
// Think of this like you ordered a drink online to pick up, and
// our coffee shop gives you a pick up code. That pick up code is passed
// from you to the kiosk stand to verify it's your order.

func CoffeeShopPickUpOrder(pickupCode int) {
	// Set a code value in our context.
	// This is the code we expect the customer to tell us to pick up an order.

	orderCode := 42069
	ctx := context.WithValue(context.Background(), "code", orderCode)

	if val := ctx.Value("code").(int); val == pickupCode {
		fmt.Println("Order picked up at kiosk! Enjoy your coffee.")
	} else {
		fmt.Printf("Wrong order number! We're looking for %d Are you stealing someone's drink?", orderCode)
	}
}
