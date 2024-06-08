package main

type Jeans struct {
	model  string
	size   int
	color  string
	length int
	price  float32
}

func (j *Jeans) accept(v Visitor) {
	v.visitForJeans(j)
}

func (j *Jeans) getType() string {
	return "jeans"
}
