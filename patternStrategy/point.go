package main

type Point struct {
	x    int
	y    int
	name string
}

func initPoint(x int, y int, name string) *Point {
	return &Point{
		x:    x,
		y:    y,
		name: name,
	}
}
