package main

import (
	"fmt"
	// go fmt puts these in alphbetical order.
)

var (
	mmm = map[string]int{
		"abc":123,
		"fff":456,
	}
)
type ss struct {
	a int
}
func main() {
	m := map[string]float64{
		"1234":3.134,
	}
	m["1"] = 1
	m["pi"] = 3.1415
	m2 := m
	m2["pi"] = 12
	fmt.Println(m, m2)

	m3 := make(map[int]ss)
	s := ss{
		a: 345,
	}
	m3[12] = s
	s2 := m3[12]
	s2.a =1111
	fmt.Println(m3)
}
