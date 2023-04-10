package main

import "fmt"

type person struct {
	name
}

func (p person) getName() string {
	return fmt.Sprintf("The person is -> %v\n", p)

}

type name string

func (n name) getName() string {
	return fmt.Sprintf("The name is -> %s\n", n)
}

func (n name) bakaboosh() {
	fmt.Println("BakaBoosh")
}

func (n name) ntbakaboosh() {
	fmt.Println("BakaBoosh")
}

func main() {
	p := person{}
	p.bakaboosh()
}
