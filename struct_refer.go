package main

import "fmt"
import "reflect"

type S struct {
	nums []int
	name string
}

func Do(arr *[]int) {
	*arr = make([]int, 4)
	(*arr)[0] = 123
}
type R struct {
	name string
	ss S
}
func main() {
	
	r := R{}
	s := S{
		name: "aaaa",
		nums: []int{3, 5},
	}
	r.ss = s
	fmt.Println(r)
	s.name = "abc"
	s.nums[0] = 999
	fmt.Println(r, s)

	rv1 := reflect.ValueOf(r)

	fmt.Println(rv1.MapKeys())
}
