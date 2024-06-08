package main

import "fmt"

type Document struct {
	name         string
	maxSize      int
	isChangeable bool
	size         int
	typeDoc      string
	sendBy       []string
}

func (d *Document) getType(typeDoc string, maxSize int, isChangeable bool) {
	d.typeDoc = typeDoc
	d.maxSize = maxSize
	d.isChangeable = isChangeable
}

func (d *Document) getDocument(size int, name string) {
	if d.maxSize < size {
		fmt.Println("error size")
		return
	}
	d.size = size
	d.name = name
}

func (d *Document) sendDocument(sendBy []string) {
	copy(d.sendBy, sendBy)
}

func (d *Document) printDocument() {
	fmt.Printf("Document '%s.%s' (%d bytes) printed\n", d.name, d.typeDoc, d.size)
}
