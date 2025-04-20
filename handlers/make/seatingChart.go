package make

import "fmt"

// In this example, we will make a slice that represents the seating chart for a coffee shop.
// The number 0 will represent a seat. You can adjust the number of seats to change the current customer count
// and increase or decrease capacity for our coffee shop

func CoffeeShopSeatingChart() {
	seatingChart := make([]int, 5, 10) // adjust these numbers to change current customer count and increase or decrease capacity

	fmt.Println("Seating chart:", seatingChart)

	// How many seats are currently available?
	fmt.Println("Seats available:", len(seatingChart))

	// What's the capacity of seats?
	fmt.Println("Capacity:", cap(seatingChart))

	// How many seats remain
	fmt.Println("Seats remaining for customers:", cap(seatingChart)-len(seatingChart))
}
