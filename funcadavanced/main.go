package main

import "fmt"

func main() {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var f = func(i []int) {
		for in, v := range i {
			if v%2 != 0 {
				i[in] = 0
			}
		}
	}
	fmt.Printf("%v\n", inout(f, i...)())

}

func inout(f func(i []int), i ...int) func() []int {
	f(i[:])
	return func() []int {
		return i
	}
}
