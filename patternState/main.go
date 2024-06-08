package main

import (
	"fmt"
	"log"
)

func main() {
	vendingMachine := newVendingMachine(1, 10, 3)

	err := vendingMachine.insertMoney(15)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.startCoffee()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.setCheck()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.setChange()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.finishWork()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()
	err = vendingMachine.insertMoney(15)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.setChange()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.finishWork()
	if err != nil {
		log.Fatalf(err.Error())
	}

}
