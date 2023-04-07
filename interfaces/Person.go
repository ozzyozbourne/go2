package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
}

func (p *person) display() {
	fmt.Printf("FirstName :%s\tLastName :%s\n", p.FirstName, p.LastName)
}
