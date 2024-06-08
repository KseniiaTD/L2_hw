package main

type Equipment interface {
	execute(*Detail)
	setNext(Equipment)
}
