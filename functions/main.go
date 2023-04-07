package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	t := []int{1, 2, 3, 4, 5, 6, 7, 8}
	p := (*reflect.SliceHeader)(unsafe.Pointer(&t))
	var a any = t
	ip := unsafe.Pointer(&a)
	fmt.Printf("%p\n", ip)

	fmt.Printf("%v\t%d\t%d\t%d\n", t, len(t), cap(t), p.Data)
	foo(t...)
	fmt.Printf("%v\t%d\t%d\t%d\n", t, len(t), cap(t), p.Data)
	a = true
	switch e := a.(type) {
	case int:
		fmt.Printf("%T\t%[1]v\n", e)
	case string:
		fmt.Printf("%T\t%[1]v\n", e)
	case bool:
		fmt.Printf("%T\t%[1]v\n", e)
	case []int:
		fmt.Printf("%T\t%[1]v\n", e)
	default:
		fmt.Printf("%s\n", "No Type Matched")
	}

}

func foo(x ...int) {
	x[0] = 1000
	p := (*reflect.SliceHeader)(unsafe.Pointer(&x))
	fmt.Printf("%v\t%d\t%d\t%d\n", x, len(x), cap(x), p.Data)
}
