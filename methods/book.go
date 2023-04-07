package main

import "fmt"

type book struct {
	title string
	price float64
}

func (b *book) print() {
	fmt.Printf("%-15s: $%-15.2f: %-15p: %-15p:  %-15p\n",
		b.title, b.price, &b.price, &b.title, b)
}

func (b *book) modify() {
	b.price -= 5
}
