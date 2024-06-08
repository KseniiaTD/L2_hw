package main

func main() {
	p1 := Point{name: "A", x: 2, y: 4}
	p2 := Point{name: "B", x: 5, y: 13}
	p3 := Point{name: "C", x: 8, y: 28}

	route := Route{pointList: []Point{p1, p2, p3}}

	road1 := Bycicle{}

	route.setRoad(&road1)
	route.buildingRoute()

	route.delPoint(&p2)
	road2 := Car{}
	route.setRoad(&road2)
	route.buildingRoute()

	p4 := Point{name: "D", x: 10, y: 20}
	route.addPoint(&p4)
	road3 := OnFoot{}
	route.setRoad(&road3)
	route.buildingRoute()
}
