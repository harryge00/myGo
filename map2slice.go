package main

import (
	"fmt"
	// go fmt puts these in alphbetical order.
)

var (
	m = map[string]int{
		"abc":123,
		"fff":456,
		"aaa":999,
	}
)

func main() {
	s := make([]int, len(m))
	fmt.Println(s, len(s), cap(s))
	for _, v := range m {
		s = append(s, v)
	}
	fmt.Println(s, len(s), cap(s))

}
