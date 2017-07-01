package main

import (
	"fmt"
)

func Fibon(n int) int {
	if n <= 1 {
		return 1
	}
	p, q := 1, 1
	for i := 2; i <= n; i++ {
		tmp := q
		q += p
		p = tmp
	}
	return q
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(Fibon(i))
	}
}
