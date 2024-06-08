package main

import "fmt"

type BuyFacade struct {
	account      Account
	pincode      PinCode
	bucket       Bucket
	address      Address
	notification Notification
}

func newBuyFacade(accountName string) *BuyFacade {
	fmt.Println("Starting create account")
	buyFacade := &BuyFacade{
		account:      *newAccount(accountName),
		pincode:      PinCode{},
		bucket:       *newBucket(),
		address:      Address{},
		notification: Notification{},
	}
	fmt.Println("Account created")
	return buyFacade
}

func (b *BuyFacade) setAddress(address *Address) {
	b.address = *address
}

func (b *BuyFacade) setPinCode(pin *PinCode) {
	b.pincode = *pin
}

func (b *BuyFacade) setNotification(msg *Notification) {
	b.notification = *msg
}

func (b *BuyFacade) getBucket() *Bucket {
	return &b.bucket
}

func (b *BuyFacade) buyProducts(accountId int, pinCode int, bucketId int, addressId int) string {
	status := true
	if !b.account.checkAccount(accountId) || !b.pincode.check(pinCode) || !b.address.checkAddress(addressId) {
		status = false
	}

	order, err := b.bucket.getOrder(bucketId)
	if err != nil {
		status = false
	}

	name := b.account.getName()

	addressStr := b.address.getAddress()
	order += b.address.getPrice()

	b.notification = *newNotification(name, order, addressStr, status)

	return b.notification.getMsg()
}
