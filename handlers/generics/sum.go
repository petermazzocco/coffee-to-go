package generics

import (
	"fmt"
	"time"
)

// Imagine we have two customers that need a bill
// But our POS can only handle one type at a time (float vs int)
// meaning if a customer paid $4.99, then we can only process orders
// that have float values with it, and vice versa for int. That
// doesn't seem efficient, right?

func sumIntPrices(m map[string]int64) int64 {
	var sum int64
	for _, v := range m {
		sum += v
	}
	return sum
}

func sumFloatPrices(m map[string]float64) float64 {
	var sum float64
	for _, v := range m {
		sum += v
	}
	return sum
}

// You can replace "int64 | float64" below with this type for better readability
type Numbers interface {
	int64 | float64
}

// With generics, we can write a single function that can handle both float and int types
// by using a type parameter V that can be either int64 or float64
// (optionally, you can create an interface for better type safety, see above)
// and by using a type parameter K that can be any comparable type
// now we can use one POS to handle all types of prices
func sumPrices[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func GiveCustomerBill() {
	floatPrices := map[string]float64{
		"Coffee": 4.99,
		"Tea":    2.99,
	}
	intPrices := map[string]int64{
		"Coffee": 4,
		"Tea":    2,
	}
	fmt.Printf("This is the example of if we can only do one type per bill on our POS\n")
	fmt.Printf("FLOAT: Customer 1 is paying a total of $%f\n", sumFloatPrices(floatPrices))
	fmt.Printf("INT: Customer 2 is paying a total of $%d\n\n", sumIntPrices(intPrices))

	time.Sleep(2 * time.Second)
	fmt.Printf("This is the example of if we can handle all COMPARABLE types on our POS \n")
	fmt.Printf("FLOAT: Customer 1 is paying a total of $%f\n", sumPrices(floatPrices))
	fmt.Printf("INT: Customer 2 is paying a total of $%d\n", sumPrices(intPrices))
}
