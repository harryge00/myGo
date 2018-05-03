package main

import (
	"fmt"
)

// func putZeroToEnd(arr []int) []int {
//  lo, hi := 0, len(arr) -1

//  for lo < hi {
//   for lo < hi && arr[lo] != 0 {
//     lo++
//   }
//   for lo < hi && arr[hi] == 0 {
//     hi --
//   }
//   if arr[lo] == 0 && arr[hi] != 0 {
//     fmt.Println("find zero", lo, hi)
//     arr[lo] = arr[hi]
//     arr[hi] = 0
//   }
//  }
//  return arr
// }

func putZeroToEnd(arr []int) []int {
 lastNonZeroNext := 0
 zeroCount := 0
 for i := 0; i < len(arr); i++ {
  if arr[i] != 0 {
    if i > lastNonZeroNext {
      arr[lastNonZeroNext] = arr[i]
      // arr[firstZero] = 0
      // firstZero--
    }
    lastNonZeroNext++
  } else {
    zeroCount++
  }
 }
 for i := 0; i < zeroCount; i++ {
  arr[len(arr) - i - 1] = 0
 }
 return arr
}

func moveZero(arr []int) []int {
    n := len(arr)
    var m int
    for i := 0; i < n; i++ {
        if (arr[i] != 0) {
           arr[m] = arr[i]
           m++
        }
    }
    
    for ;m < n; m++ {
        arr[m] = 0
    }
    
    return arr
}
// 20:46 discuss
//

func main() {
	fmt.Println(moveZero([]int{2, 5, 0, 10, 0, 15, 20}))
	fmt.Println(moveZero([]int{0, 2, 5, 0, 10, 15, 20, 0}))

}
