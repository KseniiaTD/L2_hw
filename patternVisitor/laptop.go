package main

type Laptop struct {
	model    string
	cpu      string
	diagonal float32
	color    string
	diskType string
	diskGb   int
	os       string
	price    float32
}

func (l *Laptop) accept(v Visitor) {
	v.visitForLaptop(l)
}

func (l *Laptop) getType() string {
	return "laptop"
}
