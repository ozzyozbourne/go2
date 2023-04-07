package main

import "fmt"

type books struct {
	Title  string
	Author string
}

func (b *books) display() {
	fmt.Printf("Title :%s\tAuthor :%s\n", b.Title, b.Author)
}
