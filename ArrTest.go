package main

import (
	"fmt"
)

func main() {
	var arr []int
	arr = make([]int, 1, 100)
	arr[0] = 9
	fmt.Println(arr[0])
	fmt.Println(len(arr), cap(arr))
	arr[2] = 99
}
