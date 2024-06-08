package main

import (
	"fmt"
	"math/rand"
	"time"
)

type PinCode struct {
	pin        int
	createDate time.Time
}

func newPinCode() *PinCode {
	fmt.Println("Starting create pinCode")
	pinCode := &PinCode{
		pin:        rand.Intn(9999),
		createDate: time.Now(),
	}
	fmt.Println("PinCode created")
	return pinCode
}

func (p *PinCode) check(pin int) bool {
	return pin == p.pin && time.Now().Before(p.createDate.Add(5*time.Minute))
}
