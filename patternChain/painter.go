package main

import "fmt"

type Painter struct {
	next Equipment
}

func (p *Painter) execute(d *Detail) {
	if d.isPaintDone {
		fmt.Println("Paint is already done")
		p.next.execute(d)
		return
	}
	if d.unweldedPartsCnt != 0 {
		fmt.Println("Several details are not welded")
		d.discontinued = true
		return
	}

	fmt.Println("Start paint")
	d.isPaintDone = true
	d.evenlyColored = true
	fmt.Println("End paint")
	p.next.execute(d)
}

func (p *Painter) setNext(next Equipment) {
	p.next = next
}
