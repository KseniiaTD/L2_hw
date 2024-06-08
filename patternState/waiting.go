package main

import "fmt"

type Waiting struct {
	vendingMachine *VendingMachine
}

func (i *Waiting) insertMoney(money int) error {
	fmt.Println("contact technical support number")
	return fmt.Errorf("vending Machine is broken")
}

func (i *Waiting) startCoffee() error {
	fmt.Println("contact technical support number")
	return fmt.Errorf("vending Machine is broken")
}
func (i *Waiting) setChange() error {
	fmt.Println("contact technical support number")
	return fmt.Errorf("vending Machine is broken")
}
func (i *Waiting) setCheck() error {
	fmt.Println("contact technical support number")
	return fmt.Errorf("vending Machine is broken")
}
func (i *Waiting) finishWork() error {
	if i.vendingMachine.isBroken {
		fmt.Println("contact technical support number")
		return fmt.Errorf("vending Machine is broken")
	}
	fmt.Println("Have a good day!")
	i.vendingMachine.setState(i.vendingMachine.gettingMoney)
	return nil
}
