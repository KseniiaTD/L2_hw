package main

type VendingMachine struct {
	gettingMoney  State
	gettingChange State
	makingCoffee  State
	printingCheck State
	waiting       State

	currentState State

	capsuleCoffeeCnt int
	coffeePrice      int
	change           int
	checkPaper       int
	isBroken         bool
}

func newVendingMachine(capsuleCoffeeCnt int, coffeePrice int, checkPaper int) *VendingMachine {
	v := VendingMachine{
		capsuleCoffeeCnt: capsuleCoffeeCnt,
		coffeePrice:      coffeePrice,
		checkPaper:       checkPaper,
	}
	gettingMoneyVal := GettingMoney{vendingMachine: &v}
	makingCoffeeVal := MakingCoffee{vendingMachine: &v}
	printingCheckVal := PrintingCheck{vendingMachine: &v}
	gettingChangeVal := GettingChange{vendingMachine: &v}
	waitingVal := Waiting{vendingMachine: &v}

	v.setState(&gettingMoneyVal)
	v.gettingMoney = &gettingMoneyVal
	v.makingCoffee = &makingCoffeeVal
	v.printingCheck = &printingCheckVal
	v.gettingChange = &gettingChangeVal
	v.waiting = &waitingVal
	return &v
}

func (v *VendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *VendingMachine) startCoffee() error {
	return v.currentState.startCoffee()
}

func (v *VendingMachine) setChange() error {
	return v.currentState.setChange()
}

func (v *VendingMachine) setCheck() error {
	return v.currentState.setCheck()
}

func (v *VendingMachine) finishWork() error {
	return v.currentState.finishWork()
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

func (v *VendingMachine) decrementCoffeeCnt(count int) {
	v.capsuleCoffeeCnt -= count
}
