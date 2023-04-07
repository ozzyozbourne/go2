package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Printf

type v interface {
	value()
}

type p interface {
	pointer()
}

type person struct {
	string
}

func (p *person) pointer() {
	p.string = "John Doe"
	pl("From Pointer -> %s\n", p.string)
}

func (p person) value() {
	p.string = "Messi"
	pl("From Value -> %s\n", p.string)
}

func main() {
	p3 := person{"ozzy"}
	p3.value()
	pl("Original -> %s\n", p3.string)
	p3.pointer()
	pl("Original -> %s\n", p3.string)

	pl("%s\n", strings.Repeat(" -*- ", 30))

	p1 := person{"Miccky the mouse"}
	pl("Org -> %s\n", p1.string)
	p1.pointer()
	pl("Org -> %s\n", p1.string)

	pl("\n%s\n\n", strings.Repeat(" -*- ", 30))

	p2 := person{"The New World Order !!!!!"}
	var i v = p2
	pl("Org -> %+v\n", i)
	i.value()
	pl("Org -> %+v\n", i)

	var j p = &p2
	pl("Org -> %+v\n", j)
	j.pointer()
	pl("Org -> %+v\n", j)

	pl("\n%s\n\n", strings.Repeat(" -*- ", 30))

	p4 := person{"Babadook"}
	var ij v = &p4
	ij.value()
	pl("%+v\n", ij)

	fn, ok := ij.(p)
	pl("%t\n", ok)
	fn.pointer()
	pl("%+v\n", ij)

	pl("\n%s\n\n", strings.Repeat(" -*- ", 30))
	p5 := &person{"skdckc"}
	pl("Org -> %+v\n", p5)
	p5.value()
	pl("Org -> %+v\n", p5)

}
