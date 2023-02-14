// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"fmt"
)

func main() {
	sss := make([]int, 123)
	fmt.Println(len(sss))
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

	concateSlice := append(slice[0:2], 9999, slice[2:]...)
	fmt.Println("concateSlice", concateSlice)
	fmt.Println(slice[3:], slice[0:3])

	threeSlice := []int{9, 9, 9}
	fmt.Println(append(slice, threeSlice...))

	initSlice := make([]int, 9)
	fmt.Println(initSlice)
	initSlice = append(initSlice, 123)
	initSlice = append(initSlice, 456)
	fmt.Println(initSlice)
}
