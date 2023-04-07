package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	fmt.Printf("%-20d%-40[1]b\n", kb)
	fmt.Printf("%-20d%-40[1]b\n", mb)
	fmt.Printf("%-20d%-40[1]b\n", gb)

}
