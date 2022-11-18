package main

import (
	"fmt"
)

type Order struct {
	ordId        int    `json:"order_id" validate:"required"`
	customerName string `json:"customerName" validate:"required"`
	callback     func() `json:"call_back" validate:"required"`
}

func main() {
	o := Order{
		ordId:        100,
		customerName: "Jack",
		callback:     func() {},
	}
	InterfaceToStruct(o)
}
func InterfaceToStruct(object interface{}) {
	obj := object.(Order)

	fmt.Printf("Object = %v\n", obj)
}
