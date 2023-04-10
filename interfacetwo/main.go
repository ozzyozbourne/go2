package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type ByteCounter int

type person struct {
	string
	int
}

func (t TestSort) Len() int {
	return len(t)
}

func (t TestSort) Less(i, j int) bool {
	return t[i] > t[j]
}

func (t TestSort) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type TestSort []int

func (b *ByteCounter) Write(p []byte) (n int, err error) {
	*b += ByteCounter(len(p))
	return len(p), err
}

func (p *person) String() string {
	return fmt.Sprintf("Hi My name is %s My Age is %d", p.string, p.int)
}

func main() {
	var b ByteCounter
	var w io.Writer = &b
	w.Write([]byte("Hello World"))
	fmt.Println(b)
	_, err := fmt.Fprintf(&b, "Hello World")
	if err != nil {
		return
	}
	fmt.Println(b)

	u := person{"osaid", 29}
	u2 := new(person)
	fmt.Printf("%s\n%s\n", &u, u2)
	ap := new(strings.Builder)
	ap.WriteString("slkdn")
	ap.WriteString("\nskdnvlskndv")
	fmt.Println(ap.String())

	t := TestSort{5, 6, 3, 5, 6, 8, 2, 4, 8, 56, 4, 6, 4, 33, 5, 7, 7, 3, 5, 677, 4, 3, 56, 7}

	sort.Sort(t)
	fmt.Printf("%v\n", t)
}
