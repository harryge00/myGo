package main

import (
	"fmt"
)

func MultiplyArray(arr1 []float64, num float64) {
	for i := range arr1 {
		arr1[i] *= num
	}
}


func main() {
	var arr1 []int
	fmt.Println(arr1 == nil)
	arr2 := append(arr1, 10)
	fmt.Println(arr1, arr2)
	arr3 := arr2
	arr3[0] = 99
	fmt.Println(arr2, arr3)
	arr3 = append(arr3, 1234)
	fmt.Println(arr2, arr3)
	var testmake []int
	fmt.Println(testmake == nil)

	testmake = make([]int, 0)
	fmt.Println(testmake == nil)
	fmt.Println(len(testmake), cap(testmake))
}
