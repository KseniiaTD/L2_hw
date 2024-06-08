package main

type roadsterBuilder struct {
	isRoofOpened bool
	doors        int
	trunkAmount  float32
	isTruncOpend bool
	drive        int
}

func (b *roadsterBuilder) setRoof() {
	b.isRoofOpened = true
}
func (b *roadsterBuilder) setDoors() {
	b.doors = 2
}
func (b *roadsterBuilder) setTrunk() {
	b.trunkAmount = 100
	b.isTruncOpend = false
}
func (b *roadsterBuilder) setDrive() {
	b.drive = 400
}

func (b *roadsterBuilder) getCar() car {
	return car{
		isRoofOpened: b.isRoofOpened,
		doors:        b.doors,
		trunkAmount:  b.trunkAmount,
		isTruncOpend: b.isTruncOpend,
		drive:        b.drive,
	}
}
