// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"fmt"
)

func main() {
	var a [4]int
	a[0] = 5
	var b []int
	fmt.Printf("%d %d %d\n", len(a), a[0], len(b))
	m := make(map[string]int)
	fmt.Println(m["abc"])
	m["abc"] += 1
	fmt.Println(m["abc"])
}
