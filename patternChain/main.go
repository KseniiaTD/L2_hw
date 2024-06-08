package main

import "fmt"

func main() {
	assembler := Assembler{}

	painter := Painter{}
	painter.setNext(&assembler)

	welder := Welder{}
	welder.setNext(&painter)

	stamper := Stamper{}
	stamper.setNext(&welder)

	detail1 := Detail{name: "car", unweldedPartsCnt: 10, isDefective: false}
	detail2 := Detail{name: "laptop", unweldedPartsCnt: 3, isDefective: true}

	stamper.execute(&detail1)
	if detail1.discontinued {
		fmt.Printf("Detail '%s' is broken\n", detail1.name)
	}
	stamper.execute(&detail2)
	if detail2.discontinued {
		fmt.Printf("Detail '%s' is broken\n", detail2.name)
	}
}
