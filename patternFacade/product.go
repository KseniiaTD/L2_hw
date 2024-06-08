package main

import (
	"fmt"
	"math/rand"
)

type Product struct {
	productId    int
	productName  string
	productPrice float32
}

func newProduct(name string, price float32) *Product {
	fmt.Println("Starting create product")
	product := &Product{
		productId:    rand.Int(),
		productName:  name,
		productPrice: price,
	}
	fmt.Println("Product created")
	return product
}
