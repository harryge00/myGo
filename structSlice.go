package main

import "fmt"

type One struct {
	nums []int
	name string
}

type Slice struct {
	items []One
}

func main() {

	s1 := One{
		name: "aaaa",
		nums: []int{3, 5},
	}
	s2 := One{
		name: "bbbb",
		nums: []int{3, 5, 7},
	}
	s3 := One{
		name: "cccc",
		nums: []int{3, 5, 7, 9},
	}

	sl := Slice{
		items: []One{s1, s2, s3},
	}
	fmt.Println(sl)
	x := sl.items[0]
	x.name = "xxxx"
	fmt.Println(sl)
	fmt.Println(x)
	x2 := x
	x2.name = "x2x2"
	fmt.Println(x, x2)

	var testS []*One
	for i := range sl.items {
		testS = append(testS, &sl.items[i])
	}
	fmt.Println(testS)
}
