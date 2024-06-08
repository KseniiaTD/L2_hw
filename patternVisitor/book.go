package main

type Book struct {
	name            string
	author          string
	publishingHouse string
	binding         string
	pageCnt         int
	price           float32
}

func (b *Book) accept(v Visitor) {
	v.visitForBook(b)
}

func (b *Book) getType() string {
	return "book"
}
