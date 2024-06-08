package main

import "fmt"

type Assembler struct {
	next Equipment
}

func (a *Assembler) execute(d *Detail) {
	if d.isAssembleDone {
		fmt.Println("Assemble is already done")
		a.next.execute(d)
		return
	}
	if !d.evenlyColored {
		fmt.Println("Detail is not evenly colored")
		d.discontinued = true
		return
	}
	fmt.Println("Start assemble")
	d.isAssembleDone = true
	fmt.Println("End assemble")
	fmt.Printf("Detail '%s' is ready\n", d.name)
}

func (a *Assembler) setNext(next Equipment) {
	a.next = next
}
