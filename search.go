package main

import (
	"fmt"
)

func binSearch(arr []int, num int) int {
	lo, hi := 0, len(arr) - 1
	mid := lo + (hi - lo)/2
	for lo < hi {
		if arr[mid] == num {
			return mid
		} else if arr[mid] < num {
			lo = mid + 1
		} else {
			hi = mid -1
		}
		mid = lo + (hi - lo)/2
	}
	if(arr[mid] < num) {
		return mid + 1
	}
	return mid
}


func main() {
	fmt.Println(binSearch([]int{2, 5, 10, 15, 20}, 4))
	fmt.Println(binSearch([]int{2, 5, 10, 15, 20}, 6))
	fmt.Println(binSearch([]int{2, 5, 10, 15, 20}, 0))
	fmt.Println(binSearch([]int{2, 5, 10, 15, 20}, 25))

}
