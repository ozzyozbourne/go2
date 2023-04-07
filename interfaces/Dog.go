package main

import "fmt"

type dog struct {
	Name string
	Age  int
}

func (d dog) display() {
	fmt.Printf("Name :%s\tAge :%d\n", d.Name, d.Age)
}

func (d *dog) boo() {
	fmt.Printf("%s\n", "Iam booing YOU!!!!!!!")
}
