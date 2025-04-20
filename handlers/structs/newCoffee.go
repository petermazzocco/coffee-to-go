package structs

import "fmt"

type coffee struct {
	variant string
	roast   string
	price   string
}

func CreateNewCoffee() {
	icedCoffee := coffee{
		variant: "Iced",
		roast:   "Medium",
		price:   "2.50",
	}

	fmt.Println("Our brand new iced coffee:", icedCoffee)
}
