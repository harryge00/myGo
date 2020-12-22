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
	slice := []int{1, 2, 3, 4, 5}
	emptySlice := slice[:0]
	oneSlice := slice[:1]
	twoSlice := slice[2:]
	fmt.Println(emptySlice, len(emptySlice), cap(emptySlice))
	fmt.Println(oneSlice, len(oneSlice), cap(oneSlice))
	fmt.Println(twoSlice, len(twoSlice), cap(twoSlice))

	fmt.Println(slice[3:], slice[0:3])
	// fmt.Println(slice[4:1])
	// fmt.Println(slice[4:-1])
}
