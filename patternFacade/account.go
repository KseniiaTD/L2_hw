package main

import (
	"fmt"
	"math/rand"
)

type Account struct {
	name      string
	accountId int
	isBlocked bool
}

func newAccount(accountName string) *Account {
	fmt.Println("Starting create account")
	account := &Account{
		name:      accountName,
		accountId: rand.Int(),
		isBlocked: false,
	}
	fmt.Println("Account created")
	return account
}

func (acc *Account) getName() string {
	return acc.name
}

func (acc *Account) checkAccount(accountId int) bool {
	return !acc.isBlocked && accountId == acc.accountId
}
