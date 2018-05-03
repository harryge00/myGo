package main

import (
	"fmt"
)
type T struct{
	a int
}
func f() (tp *T) {
	tp = &T{
		a: 123,
	}
	back := tp
	fmt.Println(tp)
	tp, b := &T{
		a: 456,
	}, 99
	fmt.Println(tp, b)
	fmt.Println(back)
	return tp
}

func main() {
	fmt.Println(f())
	var arr []int
	fmt.Println(arr == nil)
	arr = append(arr, 12)
	
}
