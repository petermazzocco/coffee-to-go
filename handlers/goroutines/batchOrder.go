package goroutines

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
}

func barista(ctx context.Context, wg *sync.WaitGroup, id int, orderQueue <-chan int, results chan<- Order) {
	defer wg.Done()
	for {
		select {
		// the orderQueue receives the order ID
		case o := <-orderQueue:
			workTime := time.Duration(50+rand.Intn(100)) * time.Millisecond
			log.Printf("Barista #%d is starting to work on order #%d", id, o)
			time.Sleep(workTime)
			log.Printf("Barista #%d took %v to complete order #%d", id, workTime, o)
			results <- Order{ID: o, Status: "fulfilled"}
		case <-ctx.Done():
			return
		}
	}
}

func CoffeeShopBatchOrder() {
	ctx, cancel := context.WithCancel(context.Background())
	numBaristas := 4    // adjust the number of workers to complete the task
	start := time.Now() // track the time to complete the work

	totalOrders := rand.Intn(200) + 100 // adjust the amount of random orders
	orderQueue := make(chan int, totalOrders)
	completedOrders := make(chan Order, totalOrders)

	log.Printf("The Coffee Shop just received %d orders", totalOrders)

	var wg sync.WaitGroup
	// For each barsita we will run the goroutine
	for i := range numBaristas {
		wg.Add(1)
		go barista(ctx, &wg, i, orderQueue, completedOrders)
	}

	// Send each order to the orderQueue channel
	for o := 1; o <= totalOrders; o++ {
		orderQueue <- o
	}

	// Wait and close the orderQueue channel
	go func() {
		wg.Wait()
		close(orderQueue)
	}()

	completed := make([]Order, 0, totalOrders)
	for i := 0; i < totalOrders; i++ {
		completed = append(completed, <-completedOrders)
	}

	// Wait and close the completedOrders
	go func() {
		wg.Wait()
		close(completedOrders)
	}()

	cancel()

	log.Printf("Completed %d orders in %v seconds", len(completed), time.Since(start))
}
