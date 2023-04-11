package main

import (
	"fmt"
	"sort"
)

type myIntArr []int

func (m myIntArr) Len() int {
	return len(m)
}
func (m myIntArr) Less(gIndex, sIndex int) bool {
	return m[gIndex] > m[sIndex]
}
func (m myIntArr) Swap(gIndex, sIndex int) {
	m[gIndex], m[sIndex] = m[sIndex], m[gIndex]
}

func main() {
	a := []int{1, 2}
	sort.Sort(myIntArr(a))
	fmt.Printf("%v\n", a)
}

// Go do something
func adder(i ...int) (sum int) {
	for _, v := range i {
		sum += v
	}
	return
}
