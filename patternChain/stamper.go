package main

import "fmt"

type Stamper struct {
	next Equipment
}

func (s *Stamper) execute(d *Detail) {
	if d.isStampDone {
		fmt.Println("Stamp is already done")
		s.next.execute(d)
		return
	}
	if d.isDefective {
		fmt.Println("Detail has defects")
		d.discontinued = true
		return
	}
	fmt.Println("Start stamp")
	d.isStampDone = true
	for i := d.unweldedPartsCnt; i > 0; i-- {
		d.unweldedPartsCnt--
	}
	fmt.Println("End stamp")
	s.next.execute(d)
}

func (s *Stamper) setNext(next Equipment) {
	s.next = next
}
