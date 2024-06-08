package main

import "fmt"

type Notification struct {
	message string
}

func newNotification(name string, order float32, address string, status bool) *Notification {
	var message string
	if !status {
		message = "Payment failed. Please try again"
	} else {

		message = fmt.Sprintf("Dear %s! You bought products for the %f dollars. Products will be delivered to your address: %s", name, order, address)
	}
	notification := Notification{
		message: message,
	}

	return &notification
}

func (n *Notification) getMsg() string {
	return n.message
}
