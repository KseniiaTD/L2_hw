package main

import "fmt"

type GettingChange struct {
	vendingMachine *VendingMachine
}

func (i *GettingChange) insertMoney(money int) error {

	return fmt.Errorf("Error")
}

func (i *GettingChange) startCoffee() error {
	return fmt.Errorf("Error")
}
func (i *GettingChange) setChange() error {
	if i.vendingMachine.change > 0 {
		fmt.Printf("Please, take your change: %d rub\n", i.vendingMachine.change)
		i.vendingMachine.change = 0
		i.vendingMachine.setState(i.vendingMachine.waiting)
	} else {
		i.vendingMachine.setState(i.vendingMachine.waiting)
	}
	return nil
}
func (i *GettingChange) setCheck() error {
	return fmt.Errorf("Error")
}
func (i *GettingChange) finishWork() error {
	return fmt.Errorf("Error")
}
