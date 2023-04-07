package main

import "fmt"

func main() {
	mb := book{"moby dick", 10}
	//mb.print()
	//mb.print()
	//(*book).print(&mb)
	//mb.modify()
	//mb.print()
	//mb.print()
	//(*book).print(&mb)
	ptr := &mb
	(*book).print(ptr)
	fmt.Printf("%p\n", &mb)
	fmt.Printf("%p\n", &mb)
	fmt.Printf("%p\n", &mb)
	bookPtr(&mb)
	bookPtr(&mb)
	bookPtr(&mb)
	a()
	fmt.Println("REturned")
	b()
}

func bookPtr(b *book) {
	fmt.Printf("%p\n", b)
}
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	fmt.Println("Returning")
	return
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}
