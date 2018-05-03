package main

import (
	"fmt"
)
type S string
func main() {
	fmt.Println("Hello, 15-440")
	var s S
	s = S("abcdef")
	fmt.Println(s)
	arr := []int{}
	fmt.Println(arr[3])
}
