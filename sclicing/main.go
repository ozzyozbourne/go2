package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%v\n", a)
	//remove from top
	fmt.Printf("%v\n", a[1:])
	//remove from bottom
	fmt.Printf("%v\n", a[:len(a)-1])
	//remove anywhere
	var p = 3
	fmt.Printf("%v\n", append(a[:p], a[p+1:]...))
	//append any element
	a = append(a[:3], 10)
	a = append(a[:])

	for i, v := range "øùþÿ" {
		fmt.Printf("Index -> %-3dValue -> %#-12Ulen -> %d\n", i, v, len("øùþÿ"))
	}
	fmt.Printf("%s\n", strings.Repeat(" -*- ", 10))

	for i, v := range []byte("øùþÿ") {
		fmt.Printf("Index -> %d\tValue -> %#U\n", i, v)
	}

}

func slicer(t []any) {

}
