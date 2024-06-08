package main

import "fmt"

type Welder struct {
	next Equipment
}

func (w *Welder) execute(d *Detail) {
	if d.isWeldDone {
		fmt.Println("Weld is already done")
		w.next.execute(d)
		return
	}
	fmt.Println("Start weld")
	d.isWeldDone = true
	fmt.Println("End weld")
	w.next.execute(d)
}

func (w *Welder) setNext(next Equipment) {
	w.next = next
}
