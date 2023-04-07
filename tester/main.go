package main

import (
	"fmt"
	"reflect"
)

type T struct {
	string
}

type PointerMethodCaller interface {
	pointerMethod()
}

type ValueMethodCaller interface {
	valueMethod()
}

func callValueMethodOnInterface(v ValueMethodCaller) {
	v.valueMethod()
}

func callPointerMethodOnInterface(p PointerMethodCaller) {
	p.pointerMethod()
}

func (t T) valueMethod() {
	fmt.Printf("%#v \t %p\n", t, &t)
}

func (t *T) pointerMethod() {
	fmt.Printf("%#v \t %p\n", t, t)
}

func main() {
	t := T{"lksvsdkvm"}
	tp := &t
	var i PointerMethodCaller = tp

	t.string = "cheap thrills"
	fmt.Printf("%s\t%s\t%s\n", t.string, reflect.ValueOf(i), tp)
	//fmt.Printf("Value created \t%#v with address %p\n", t, &t)
	//fmt.Printf("Pointer created on \t%#v with address %p\n", *tp, tp)
	//
	//
	//t.valueMethod()
	//t.pointerMethod()
	//callValueMethodOnInterface(t)
	//callPointerMethodOnInterface(tp)
	//callValueMethodOnInterface(tp)
}
