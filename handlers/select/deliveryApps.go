package selectPkg

import (
	"fmt"
	"time"
)

func DeliveryApps() {
	doordash := make(chan string)
	uberEats := make(chan string)

	fmt.Println("Waiting for orders...")

	go func() {
		time.Sleep(2 * time.Second)
		doordash <- "Doordash"
		close(doordash)
	}()

	go func() {
		time.Sleep(4 * time.Second)
		uberEats <- "Uber Eats"
		close(uberEats)
	}()

	for i := 0; i < 2; i++ {
		select {
		case dd := <-doordash:
			fmt.Printf("Received %s\n", dd)
		case ue := <-uberEats:
			fmt.Printf("Received %s\n", ue)
		}
	}
}
