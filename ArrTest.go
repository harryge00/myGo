package main

import (
	"fmt"
	"unsafe"
	"strconv"
)

func MultiplyArray(arr1 []float64, num float64) {
	for i := range arr1 {
		arr1[i] *= num
	}
}

func findMissingRanges(nums []int, lower int, upper int) []string {
    res := make([]string, 0)
    if len(nums) == 0 {
        return append(res, convert2Str(lower, upper))
    }
    if lower < nums[0] {
    	res = append(res, convert2Str(lower, nums[0] - 1))
    }
    for lastExist := nums[0], i := 1; i < len(nums); i++ {
    	if 
    }
    lastExist := nums[0]


    return res
}

func convert2Str(a, b int) string {
	if a == b {
		return strconv.Itoa(a)
	} else {
		return strconv.Itoa(a) + "->" + strconv.Itoa(b)
	}
}

func main() {
	var testarr []int
	testarr = append(testarr, 10)
	fmt.Println(testarr)
	arrf := []float64{4,5,6, 7,8,9, 10}
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
	arr3 := []int{4,5,6, 7,8,9, 10}
	fmt.Println(arr3, len(arr3), cap(arr3))
	arr4 := arr3[4:6]
	fmt.Println(arr4, len(arr4), cap(arr4))

	arr4[0] = 9999
	fmt.Println(arr3, arr4)

	m := make(map[string][]string)
	m["abc"] = append(m["abc"], "def")
	m["abc"] = append(m["abc"], "defddd")
	m["def"] = append(m["def"], "xxxfddd")
	fmt.Println(m)
	var ddd []int
	ddd = []int{2, 5, 5}
	eee := []int{3,4,5}
	fmt.Println(ddd)
	fmt.Printf("%p %p \n", &ddd, &eee)
	ddd2 := ddd[1:]
	fmt.Printf("pp %p ", &ddd2)

	fmt.Println(unsafe.Pointer(&ddd2), unsafe.Pointer(&ddd2[0]), unsafe.Pointer(&ddd2[1]))

	fmt.Println(ddd, ddd2)
	ddd2[0] = 999
	fmt.Println(ddd, ddd2)
	fmt.Println(len(ddd), cap(ddd))
	fmt.Println(len(ddd2), cap(ddd2))
	ddd2 = append(ddd2, 777)
	fmt.Println(ddd2, len(ddd2), cap(ddd2))
	ddd2 = append(ddd2, 888)
	fmt.Println(ddd2, len(ddd2), cap(ddd2))
	ddd2 = append(ddd2, 999)
	fmt.Println(ddd2, len(ddd2), cap(ddd2))
	ddd2[1] = 444
	fmt.Println(ddd2, len(ddd2), cap(ddd2))
	fmt.Printf("%p \n", &ddd2)

	fmt.Println(ddd, len(ddd), cap(ddd))


}
