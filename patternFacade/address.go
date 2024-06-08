package main

import (
	"fmt"
	"math/rand"
)

type Address struct {
	addressId int
	address   string
	price     float32
	isValid   bool
}

func newAddress(addressName string) *Address {
	fmt.Println("Starting create address")
	address := &Address{
		addressId: rand.Int(),
		address:   addressName,
		price:     rand.Float32() * 10,
		isValid:   true,
	}
	fmt.Println("Address created")
	return address
}

func (a *Address) checkAddress(addressId int) bool {
	return a.isValid && addressId == a.addressId
}

func (a *Address) getPrice() float32 {
	return a.price
}

func (a *Address) getAddress() string {
	return a.address
}
