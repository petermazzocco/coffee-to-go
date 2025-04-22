package context

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
)

// Context can be tricky to learn at first, but an easy example is
// controlling timeouts. Imagine if we have a delivery service for our
// coffee, and we guaranteed a set delivery time of 5 seconds. Anything
// over 5 seconds, would trigger a refund to our customer because the
// coffee has gone cold. Run the cmd and enter any amout of time in seconds.

func CoffeeShopDelivery(seconds string) {
	sec, err := strconv.Atoi(seconds)
	if err != nil {
		fmt.Println("Invalid input type:", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 5 seconds represents the expected delivery time
	defer cancel()

	delivered := make(chan bool) // Handle the delivery status

	go func() {
		log.Println("Coffee is being delivered...")
		duration := time.Duration(sec) * time.Second
		time.Sleep(duration) // The delivery time you entered
		close(delivered)     // Close the channel to indicate delivery completion
	}()

	select {
	case <-delivered: // The customer will "receive" a delivered order
		log.Println("Coffee delivered on time! The customer is happy.")
	case <-ctx.Done(): // Handle the case in which it's been delivered late (timed out)
		log.Println("Our coffee was delivered late! We will refund the custoemr for their cold coffee:", ctx.Err())
	}
}
