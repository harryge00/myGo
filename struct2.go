package main

import "fmt"

type T struct{
	name string
}
func main() {
	t := T{name:"aaa"}
	b := t
	b.name="bbb"
	fmt.Println(t, b)
	arr := []T{t, b}
	for _, t := range arr {
		t.name = "xxx"
	}
	fmt.Println(arr)
}
