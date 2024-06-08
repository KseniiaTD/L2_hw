package main

type State interface {
	insertMoney(money int) error
	startCoffee() error
	setChange() error
	setCheck() error
	finishWork() error
}
