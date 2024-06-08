package main

import "fmt"

type Car struct {
}

func (c *Car) build(route *Route) {
	fmt.Println("Car route built:")
	for _, el := range route.pointList {
		fmt.Println(el.name)
	}
}
