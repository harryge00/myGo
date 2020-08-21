package main

import "fmt"

func main() {
	slice := make([]int, 10, 20)
	fmt.Printf("%p, %v, len: %v; cap: %v \n", slice, slice, len(slice), cap(slice))
	k2 := append(slice, 10)
	k2 = append(slice, 20)
	fmt.Printf("%p, %v, len: %v; cap: %v \n", slice, slice, len(slice), cap(slice))
	
	fmt.Println(k2)
}