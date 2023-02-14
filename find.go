package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 5, 1, 2, 6, 8, 9, 10}

	num := 7
	lo, hi := 0, len(arr)-1
	mid := lo + (hi-lo)/2
	for lo < hi {
		if arr[mid] < num {
			lo = mid + 1
		} else if arr[mid] > num {
			hi = mid - 1
		} else {
			break
		}
		mid = lo + (hi-lo)/2
		fmt.Println(lo, hi, mid)
	}
	fmt.Println(mid)
}
