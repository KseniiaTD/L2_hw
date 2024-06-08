package main

type Sbj interface {
	getType() string
	accept(Visitor)
}
