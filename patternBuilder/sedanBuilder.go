package main

type sedanBuilder struct {
	isRoofOpened bool
	doors        int
	trunkAmount  float32
	isTruncOpend bool
	drive        int
}

func (b *sedanBuilder) setRoof() {
	b.isRoofOpened = false
}
func (b *sedanBuilder) setDoors() {
	b.doors = 4
}
func (b *sedanBuilder) setTrunk() {
	b.trunkAmount = 300
	b.isTruncOpend = false
}
func (b *sedanBuilder) setDrive() {
	b.drive = 200
}

func (b *sedanBuilder) getCar() car {
	return car{
		isRoofOpened: b.isRoofOpened,
		doors:        b.doors,
		trunkAmount:  b.trunkAmount,
		isTruncOpend: b.isTruncOpend,
		drive:        b.drive,
	}
}
