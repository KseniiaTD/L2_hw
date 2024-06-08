package main

import "fmt"

func main() {

	product1 := newProduct("Apple", 5)
	product2 := newProduct("Milk", 3.5)
	product3 := newProduct("Chips", 2)

	addr1 := newAddress("Ivanovskaya street, 12")

	pin := newPinCode()

	buyEl := newBuyFacade("Oleg")

	buyEl.bucket.addProduct(product1, product2, product3)

	buyEl.setAddress(addr1)
	buyEl.setPinCode(pin)

	fmt.Println()
	msg := buyEl.buyProducts(buyEl.account.accountId, pin.pin, buyEl.bucket.bucketId, addr1.addressId)
	fmt.Println(msg)
}
