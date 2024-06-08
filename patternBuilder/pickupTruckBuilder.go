package main

type pickupTruckBuilder struct {
	isRoofOpened bool
	doors        int
	trunkAmount  float32
	isTruncOpend bool
	drive        int
}

func (b *pickupTruckBuilder) setRoof() {
	b.isRoofOpened = false
}
func (b *pickupTruckBuilder) setDoors() {
	b.doors = 4
}
func (b *pickupTruckBuilder) setTrunk() {
	b.trunkAmount = 500
	b.isTruncOpend = true
}
func (b *pickupTruckBuilder) setDrive() {
	b.drive = 500
}

func (b *pickupTruckBuilder) getCar() car {
	return car{
		isRoofOpened: b.isRoofOpened,
		doors:        b.doors,
		trunkAmount:  b.trunkAmount,
		isTruncOpend: b.isTruncOpend,
		drive:        b.drive,
	}
}
