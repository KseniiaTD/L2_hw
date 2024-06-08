package main

import "fmt"

type MakingCoffee struct {
	vendingMachine *VendingMachine
}

func (i *MakingCoffee) insertMoney(money int) error {

	return fmt.Errorf("Error")
}

func (i *MakingCoffee) startCoffee() error {
	i.vendingMachine.decrementCoffeeCnt(1)
	fmt.Println("Coffee is ready")
	i.vendingMachine.setState(i.vendingMachine.printingCheck)
	return nil
}
func (i *MakingCoffee) setChange() error {
	return fmt.Errorf("Error")
}
func (i *MakingCoffee) setCheck() error {
	return fmt.Errorf("Error")
}
func (i *MakingCoffee) finishWork() error {
	return fmt.Errorf("Error")
}
