package main

import "fmt"

func addNum(s []int, num int) {
	s = append(s, num)

}

func main() {
	var s []int
	addNum(s, 1)
	fmt.Println("should be empty slice", s)
	addNum([]int{}, 33)
}
