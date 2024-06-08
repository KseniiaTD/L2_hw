package main

type Visitor interface {
	visitForBook(*Book)
	visitForJeans(*Jeans)
	visitForLaptop(*Laptop)
}
