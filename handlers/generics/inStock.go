package generics

import (
	"log"
	"reflect"
)

type Status interface {
	bool | string
}

func setInStockStatus[S Status](i S) {
	log.Println("The item is in stock:", i, ". The type is:", reflect.TypeOf(i))
}

func GetStatus() {
	setInStockStatus("true")
	setInStockStatus(true)
}
