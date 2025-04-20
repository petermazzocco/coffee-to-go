package context

import (
	"context"
	"fmt"
)

// With context, you can pass state information to goroutines.
// Think of this like you ordered a drink online to pick up, and
// our coffee shop gives you a pick up code. That pick up code is passed
// from you to the kiosk stand to verify it's your order.

func pickUp(ctx context.Context, kiosk chan int, orderCode int) {
	select {
	case <-ctx.Done():
		fmt.Println("The kiosk is currently closed:", ctx.Err())
		// Recieved from the kiosk channel
		// Person taking the order from the counter
	case <-kiosk:
		// Grab the value from the context and check it against the inputed ordercode
		if val := ctx.Value("code").(int); val == orderCode {
			fmt.Println("Order picked up! Enjoy your coffee.")
		} else {
			fmt.Println("Wrong order number! Are you stealing someone's drink?")
		}
	default:
		fmt.Println("No order code received!")
	}
}

func CoffeeShopPickUpOrder(orderCode int) {
	// Set a code value in our context.
	// This is the code we expect the customer to tell us to pick up an order.
	ctx := context.WithValue(context.Background(), "code", 42069)

	// Create a channel for retrieving the value
	// Think of this like the pick up kiosk station
	kioskStand := make(chan int)
	// We need to close the channel to prevent a memory leak but let's defer it to the end
	// Think of this like the pick up kiosk station staying open
	// and waiting for the customer to arrive but not closing until the order is picked up
	// Feels like a waste of resources
	defer close(kioskStand)

	go pickUp(ctx, kioskStand, orderCode)

	// Send the orderCode value to the kioskStand channel
	// We're putting the order on the counter with a code
	kioskStand <- orderCode
}
