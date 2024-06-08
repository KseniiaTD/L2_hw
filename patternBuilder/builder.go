package main

type carBuilder interface {
	setRoof()
	setDoors()
	setTrunk()
	setDrive()
	getCar() car
}

func createCar(builderType string) car {
	if builderType == "pickup truck" {
		return buildCar(&pickupTruckBuilder{})
	}
	if builderType == "roadster" {
		return buildCar(&roadsterBuilder{})
	}
	if builderType == "sedan" {
		return buildCar(&sedanBuilder{})
	}
	return car{}
}

func buildCar(b carBuilder) car {
	b.setRoof()
	b.setDoors()
	b.setTrunk()
	b.setDrive()
	return b.getCar()
}
