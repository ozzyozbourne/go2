package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Printf

func main() {
	b := books{
		Title:  "Harry Potter",
		Author: "J K Rowling",
	}
	p := person{
		FirstName: "James",
		LastName:  "Bond",
	}

	d := dog{
		Name: "Max",
		Age:  78,
	}

	i := []show{&b, &p, d}

	for _, it := range i {
		it.display()
	}

	s := struct {
		string
		int
	}{"osaid", 29}

	fmt.Printf("%#v\n", s)

	var ad generic
	ad = s
	pl("%T\n", ad)
	f := "%#v\n"
	ad = d
	pl(f, d)

	ad = pl
	pl("%T\n", ad)

	ad = b
	as, vs := ad.(books)
	pl("%#v\t%t\n", as, vs)
	ad = &b
	q := ad.(show)

	pl("%#v\n", q)

	//for _, it := range i {
	//	if q, w := it.(interface{ boo() }); w {
	//		q.boo()
	//	} else {
	//		pl("%s\n", "Bad Luck")
	//	}
	//
	//}

	ad = &d
	if q, w := ad.(interface{ boo() }); w {
		q.boo()
	} else {
		pl("%s\n", "Wrong Wrong Wrong !!!!!!!")
	}

	pl("%s\n", strings.Repeat(" -*- ", 30))
	dog.display(d)
	(*dog).display(&d)

	(*dog).boo(&d)
	pl("%s\n", strings.Repeat(" -*- ", 30))
	ast := &d
	ast.display()
	ast.boo()
	d.display()
	d.boo()

}
