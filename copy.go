package main

import "fmt"

type Child struct {
	A int
	B int
}

type Parent struct {
	name string
	Child
}

func main() {
	p := &Parent{
		name: "test",
	}
	p.A = 123
	p.B = 456
	var b *Parent
	b
	s := []*Parent{
		&(*p),
		&(*p),
		&(*p),
	}
	fmt.Println(s)
}
