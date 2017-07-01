package main

import "fmt"
import "sort"

type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
	sort.Sort(s)
	str := "["
	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}
func check(obj interface{}) {
	if d, ok := obj.(Sequence); ok {
		fmt.Println("is Sequence")
		fmt.Println(d.Newest())
	}
}
func main() {
	var tt Sequence
	tt = make([]int, 5, 6)
	tt[0] = 10
	tt[1] = 9
	tt[2] = 27
	fmt.Println(tt[0])
	fmt.Println(tt.Less(1, 2))
	check(tt)
}
