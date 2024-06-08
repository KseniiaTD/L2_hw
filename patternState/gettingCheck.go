package main

import "fmt"

type PrintingCheck struct {
	vendingMachine *VendingMachine
}

func (i *PrintingCheck) insertMoney(money int) error {

	return fmt.Errorf("Error")
}

func (i *PrintingCheck) startCoffee() error {
	return fmt.Errorf("Error")
}
func (i *PrintingCheck) setChange() error {
	return fmt.Errorf("Error")
}
func (i *PrintingCheck) setCheck() error {
	if i.vendingMachine.checkPaper > 0 {
		i.vendingMachine.checkPaper--
		fmt.Println("Check is ready")
		i.vendingMachine.setState(i.vendingMachine.gettingChange)
	} else {
		fmt.Println("No paper for check")
		i.vendingMachine.isBroken = true
		i.vendingMachine.setState(i.vendingMachine.waiting)
	}
	return nil
}
func (i *PrintingCheck) finishWork() error {
	return fmt.Errorf("Error")
}
