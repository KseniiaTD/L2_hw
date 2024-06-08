package main

import "fmt"

type Bycicle struct {
}

func (b *Bycicle) build(route *Route) {
	fmt.Println("Bike route built:")
	for _, el := range route.pointList {
		fmt.Println(el.name)
	}
}
