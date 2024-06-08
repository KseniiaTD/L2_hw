package main

import (
	"fmt"
	"strings"
)

type Route struct {
	pointList []Point
	road      BuildingRoad
}

func (r *Route) addPoint(p *Point) {
	r.pointList = append(r.pointList, *p)
}

func (r *Route) delPoint(p *Point) {
	var start bool
	for i, el := range r.pointList {
		if strings.EqualFold(el.name, p.name) {
			start = true
		}
		if start {
			r.pointList = append(r.pointList[:i], r.pointList[i+1:]...)
			break
		}
	}
	if !start {
		fmt.Println("Point not found")
	}
}

func (r *Route) setRoad(m BuildingRoad) {
	r.road = m
}

func (r *Route) buildingRoute() {
	r.road.build(r)
}
