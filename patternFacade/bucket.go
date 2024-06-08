package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Bucket struct {
	bucketId int
	products []*Product
}

func newBucket() *Bucket {
	fmt.Println("Starting create bucket")
	bucket := &Bucket{
		bucketId: rand.Int(),
	}
	fmt.Println("Bucket created")
	return bucket
}

func (b *Bucket) addProduct(product ...*Product) {
	b.products = append(b.products, product...)
}

func (b *Bucket) getOrder(bucketId int) (float32, error) {
	var order float32
	if bucketId != b.bucketId {
		return 0, errors.New("not found")
	}
	for _, el := range b.products {
		order += el.productPrice
	}

	return order, nil
}
