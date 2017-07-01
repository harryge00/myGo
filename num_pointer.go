package main

import "fmt"

func main() {
	var a *int
	const three = 3
	var i int = three
	a = &i
	fmt.Printf("%d\n", i)
	fmt.Printf("a %d,%d\n", a, *a)
}
