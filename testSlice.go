package main

import "fmt"

func main() {

	slice := make(map[int]int)
	fmt.Println(len(slice))

	slice[0] = 9
	slice[1] = 8
	slice[1] = 0, false
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	slice[2] = 0, false
	fmt.Println(slice)
	s2 := make(map[int]bool)
	fmt.Println(s2[35])
	var ddd uint
	ddd = 1
	fmt.Println(ddd - 33)
}
