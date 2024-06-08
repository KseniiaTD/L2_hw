package main

import "fmt"

type PackageSbj struct {
	typePack      string
	sizePack      int
	isFragilePack bool
}

func (p *PackageSbj) visitForBook(b *Book) {
	p.typePack = "plastic bag"
	p.sizePack = 1
	fmt.Println(p)
}

func (p *PackageSbj) visitForJeans(j *Jeans) {
	p.typePack = "plastic bag"
	p.sizePack = 2
	fmt.Println(p)
}

func (p *PackageSbj) visitForLaptop(l *Laptop) {
	p.typePack = "box"
	p.sizePack = 0
	p.isFragilePack = true
	fmt.Println(p)
}
