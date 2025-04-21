package selectPkg

import (
	"fmt"
	"time"
)

func ReceivePayment(amount int) {
	validAmount := make(chan int)

	fmt.Println("Did the customer pay the right amount for their coffe?")

	go func() {
		time.Sleep(1 * time.Second)
		validAmount <- amount
		close(validAmount)
	}()

	select {
	case amt := <-validAmount:
		if amount == 5 {
			fmt.Printf("Customer paid the RIGHT amount of $%d!\n", amt)
		} else {
			fmt.Printf("Customer paid the WRONG amount of $%d!\n", amt)
		}
	}
}
