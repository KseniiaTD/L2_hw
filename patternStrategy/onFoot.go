package main

import "fmt"

type OnFoot struct {
}

func (o *OnFoot) build(route *Route) {
	fmt.Println("Walking route built:")
	for _, el := range route.pointList {
		fmt.Println(el.name)
	}
}
