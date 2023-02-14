package main

import (
	"fmt"
)

func findLargests(arr []int) (int, int) {
	largest, largest2nd := -1, -1
	for i := 0; i < len(arr); i++ {
		if largest < 0 || arr[i] > arr[largest] {
			largest2nd = largest
			largest = i
			continue
		}
		if largest2nd < 0 || arr[i] > arr[largest2nd] {
			largest2nd = i
		}
	}
	return largest, largest2nd
}

func main() {
	arr := []int{4, 5, 1, 2, 6, 8, 9, 10}
	a, b := findLargests(arr)
	fmt.Println(arr[a], arr[b])

	arr = []int{-4, 5, 1, 2, -6, -8, 9, 10, 2, 4}
	a, b = findLargests(arr)
	fmt.Println(arr[a], arr[b])

	arr = []int{10, -4, 5, 9, 1, 2, -6, -8, 2, 4}
	a, b = findLargests(arr)
	fmt.Println(arr[a], arr[b])

	arr = []int{9, 10, -4, 5, 9, 1, 2, 10, -6, -8, 2, 4}
	a, b = findLargests(arr)
	fmt.Println(arr[a], arr[b])
}
