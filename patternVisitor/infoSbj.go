package main

import "fmt"

type InfoSbj struct {
	infoMsg string
}

func (i *InfoSbj) visitForBook(b *Book) {
	i.infoMsg = fmt.Sprintf("name: %s, author:%s, publishingHouse: %s, binding: %s, pageCnt: %d, price: %f", b.name, b.author, b.publishingHouse, b.binding, b.pageCnt, b.price)
	fmt.Println(i.infoMsg)
}

func (i *InfoSbj) visitForJeans(b *Jeans) {
	i.infoMsg = fmt.Sprintf("model: %s, size:%d, color: %s, length: %d, price: %f", b.model, b.size, b.color, b.length, b.price)
	fmt.Println(i.infoMsg)
}

func (i *InfoSbj) visitForLaptop(b *Laptop) {
	i.infoMsg = fmt.Sprintf("model: %s, cpu:%s, diagonal:%f, color: %s, diskType: %s, diskGb: %d, os: %s, price: %f", b.model, b.cpu, b.diagonal, b.color, b.diskType, b.diskGb, b.os, b.price)
	fmt.Println(i.infoMsg)
}
