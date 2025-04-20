package goroutines

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
}

func logMessage(format string, args ...any) {
	prefix := fmt.Sprintf("[%s] ", time.Now().Format("15:04:05.000"))
	fmt.Printf(prefix+format+"\n", args...)
}

func barista(id int, orders <-chan int, results chan<- Order, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case orderNum, ok := <-orders:
			if !ok {
				return
			}
			workTime := time.Duration(50+rand.Intn(100)) * time.Millisecond // adjust the speed at which the workers operate
			logMessage("Barista #%d is starting to work on order #%d", id, orderNum)
			time.Sleep(workTime)
			logMessage("Barista #%d took %v to complete order #%d", id, workTime, orderNum)
			results <- Order{ID: orderNum, Status: "fulfilled"}
		case <-ctx.Done():
			return
		}
	}
}

func CoffeeShopBatchOrder() {
	ctx, cancel := context.WithCancel(context.Background())
	numBaristas := 4 // adjust the number of workers to complete the task
	start := time.Now()

	totalOrders := rand.Intn(200) + 100 // adjust the amount of random orders
	orderQueue := make(chan int, totalOrders)
	completedOrders := make(chan Order, totalOrders)

	logMessage("The Coffee Shop just received %d orders", totalOrders)

	var wg sync.WaitGroup

	for i := range numBaristas {
		wg.Add(1)
		go barista(i, orderQueue, completedOrders, &wg, ctx)
	}

	for orderNum := 1; orderNum <= totalOrders; orderNum++ {
		orderQueue <- orderNum
	}

	go func() {
		wg.Wait()
		close(orderQueue)
	}()

	for i := 0; i < totalOrders; i++ {
		<-completedOrders
	}

	go func() {
		wg.Wait()
		close(completedOrders)
	}()

	cancel()

	logMessage("All orders completed in %v.", time.Since(start))
}
