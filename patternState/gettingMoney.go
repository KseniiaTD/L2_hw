package main

import "fmt"

type GettingMoney struct {
	vendingMachine *VendingMachine
}

func (i *GettingMoney) insertMoney(money int) error {

	if money >= i.vendingMachine.coffeePrice && i.vendingMachine.capsuleCoffeeCnt > 0 {
		i.vendingMachine.change = money - i.vendingMachine.coffeePrice
		fmt.Println("Start coffee")
		i.vendingMachine.setState(i.vendingMachine.makingCoffee)
	} else if money < i.vendingMachine.coffeePrice {
		i.vendingMachine.change = money
		i.vendingMachine.setState(i.vendingMachine.gettingChange)
	} else {
		i.vendingMachine.isBroken = true
		i.vendingMachine.change = money
		i.vendingMachine.setState(i.vendingMachine.gettingChange)
	}
	return nil
}

func (i *GettingMoney) startCoffee() error {
	return fmt.Errorf("Error")
}
func (i *GettingMoney) setChange() error {
	return fmt.Errorf("Error")
}
func (i *GettingMoney) setCheck() error {
	return fmt.Errorf("Error")
}
func (i *GettingMoney) finishWork() error {
	return fmt.Errorf("Error")
}
