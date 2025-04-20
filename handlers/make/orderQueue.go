package make

import (
	"fmt"
	"time"
)

// In this example, we will make a channel for holding order queues!
// This channel will be used to queue orders for processing and pass data
// to the next goroutine for processing.

func CoffeeShopOrderQueue(totalOrders int) {
	start := time.Now()
	// First make a channel of integers to represent the amount of orders in our queue
	// We set a buffer of the length of the total orders
	orderQueue := make(chan int, totalOrders)

	// Create a simple goroutine
	go func() {
		// For each order in the totalOrders
		for order := range totalOrders {
			// Simulate a delay for processing
			time.Sleep(time.Millisecond * 10)
			fmt.Printf("Order #%d queued\n", order)
			// The order will be sent to the orderQueue channel
			orderQueue <- order
		}
		close(orderQueue) // Close the orderqueue because there's no more orders to take
	}()

	// Wait for all orders to be processed and log out

	for order := range orderQueue {
		fmt.Printf("Order #%d processed\n", order)
	}

	// Keep running the command and see how long it takes to process all orders
	// The time should vary each time you run it!
	fmt.Printf("Total time taken: %s\n", time.Since(start))
}
