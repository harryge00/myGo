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
	var testarr []int
	testarr = append(testarr, 10)
	fmt.Println(testarr)
	arrf := []float64{4, 5, 6, 7, 8, 9, 10}

	MultiplyArray(arrf, 1000)
	fmt.Println(arrf[7:7])
	arrf = append(arrf[:6], arrf[7:]...)
	fmt.Println(arrf)
	var arr []int
	arr = make([]int, 1, 100)
	arr[0] = 9
	fmt.Println(arr)
	fmt.Println(len(arr), cap(arr))
	_ = append(arr, 23)
	fmt.Println(arr)

	fmt.Println(len(arr), cap(arr))
	arr2 := arr[3:15]
	fmt.Println(arr2, len(arr2), cap(arr2))
	arr3 := []int{4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr3, len(arr3), cap(arr3))
	arr4 := arr3[4:6]
	fmt.Println(arr4, len(arr4), cap(arr4))

	m := make(map[string][]string)
	m["abc"] = append(m["abc"], "def")
	m["abc"] = append(m["abc"], "defddd")
	m["def"] = append(m["def"], "xxxfddd")
	fmt.Println(m)
	var ddd []int
	ddd = []int{2, 5, 5}
	fmt.Println(ddd)

	testCap := make([]int, 0)
	fmt.Println("default cap of a len0 slice:", cap(testCap))
}
